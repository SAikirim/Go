/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://aws.amazon.com/apache2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package software.amazon.smithy.go.codegen.integration;

import java.util.Collection;
import java.util.HashSet;
import java.util.Objects;
import java.util.Optional;
import java.util.Set;
import java.util.function.BiPredicate;
import software.amazon.smithy.model.Model;
import software.amazon.smithy.model.shapes.OperationShape;
import software.amazon.smithy.model.shapes.ServiceShape;
import software.amazon.smithy.utils.SmithyBuilder;
import software.amazon.smithy.utils.ToSmithyBuilder;

/**
 * Represents a runtime plugin for a client that hooks into various aspects
 * of Go code generation, including adding configuration settings
 * to clients and middleware plugins to both clients and commands.
 *
 * <p>These runtime client plugins are registered through the
 * {@link GoIntegration} SPI and applied to the code generator at
 * build-time.
 */
public final class RuntimeClientPlugin implements ToSmithyBuilder<RuntimeClientPlugin> {
    private final BiPredicate<Model, ServiceShape> servicePredicate;
    private final OperationPredicate operationPredicate;
    private final Set<ConfigField> configFields;
    private final Set<ConfigFieldResolver> configFieldResolvers;
    private final MiddlewareRegistrar registerMiddleware;

    private RuntimeClientPlugin(Builder builder) {
        operationPredicate = builder.operationPredicate;
        servicePredicate = builder.servicePredicate;
        configFields = builder.configFields;
        registerMiddleware = builder.registerMiddleware;
        configFieldResolvers = builder.configFieldResolvers;
    }


    @FunctionalInterface
    public interface OperationPredicate {
        /**
         * Tests if middleware is applied to an individual operation.
         *
         * @param model     Model the operation belongs to.
         * @param service   Service the operation belongs to.
         * @param operation Operation to test.
         * @return Returns true if middleware should be applied to the operation.
         */
        boolean test(Model model, ServiceShape service, OperationShape operation);
    }

    /**
     * Gets the config fields that will be added to the client config by this plugin.
     * @return the config field resolvers.
     */
    public Set<ConfigFieldResolver> getConfigFieldResolvers() {
        return configFieldResolvers;
    }

    /**
     * Gets the optionally present middleware registrar object that resolves to middleware registering function.
     *
     * @return Returns the optionally present MiddlewareRegistrar object.
     */
    public Optional<MiddlewareRegistrar> registerMiddleware() {
        return Optional.ofNullable(registerMiddleware);
    }

    /**
     * Returns true if this plugin applies to the given service.
     *
     * <p>By default, a plugin applies to all services but not to specific
     * commands. You an configure a plugin to apply only to a subset of
     * services (for example, only apply to a known service or a service
     * with specific traits) or to no services at all (for example, if
     * the plugin is meant to by command-specific and not on every
     * command executed by the service).
     *
     * @param model   The model the service belongs to.
     * @param service Service shape to test against.
     * @return Returns true if the plugin is applied to the given service.
     * @see #matchesOperation(Model, ServiceShape, OperationShape)
     */
    public boolean matchesService(Model model, ServiceShape service) {
        return servicePredicate.test(model, service);
    }

    /**
     * Returns true if this plugin applies to the given operation.
     *
     * @param model     Model the operation belongs to.
     * @param service   Service the operation belongs to.
     * @param operation Operation to test against.
     * @return Returns true if the plugin is applied to the given operation.
     * @see #matchesService(Model, ServiceShape)
     */
    public boolean matchesOperation(Model model, ServiceShape service, OperationShape operation) {
        return operationPredicate.test(model, service, operation);
    }

    /**
     * Gets the config fields that will be added to the client config by this plugin.
     *
     * <p>Each config field will be added to the client's Config object and will
     * result in a corresponding getter method being added to the config. E.g.:
     * <p>
     * type ClientOptions struct {
     * // My docs.
     * MyField string
     * }
     * <p>
     * func (o ClientOptions) GetMyField() string {
     * return o.MyField
     * }
     *
     * @return Returns the config fields to add to the client config.
     */
    public Set<ConfigField> getConfigFields() {
        return configFields;
    }

    public static Builder builder() {
        return new Builder();
    }

    @Override
    public SmithyBuilder<RuntimeClientPlugin> toBuilder() {
        return builder()
                .configFieldResolvers(configFieldResolvers)
                .servicePredicate(servicePredicate)
                .operationPredicate(operationPredicate)
                .registerMiddleware(registerMiddleware);
    }

    /**
     * Builds a {@code RuntimeClientPlugin}.
     */
    public static final class Builder implements SmithyBuilder<RuntimeClientPlugin> {
        private BiPredicate<Model, ServiceShape> servicePredicate = (model, service) -> true;
        private OperationPredicate operationPredicate = (model, service, operation) -> false;
        private Set<ConfigField> configFields = new HashSet<>();
        private Set<ConfigFieldResolver> configFieldResolvers = new HashSet<>();
        private MiddlewareRegistrar registerMiddleware;

        @Override
        public RuntimeClientPlugin build() {
            return new RuntimeClientPlugin(this);
        }

        /**
         * Registers middleware into the operation middleware stack.
         *
         * @param registerMiddleware resolved middleware registrar to set.
         * @return Returns the builder.
         */
        public Builder registerMiddleware(MiddlewareRegistrar registerMiddleware) {
            this.registerMiddleware = registerMiddleware;
            return this;
        }

        /**
         * Sets a predicate that determines if the plugin applies to a
         * specific operation.
         *
         * <p>When this method is called, the {@code servicePredicate} is
         * automatically configured to return false for every service.
         *
         * <p>By default, a plugin applies globally to a service, which thereby
         * applies to every operation when the middleware stack is copied.
         *
         * @param operationPredicate Operation matching predicate.
         * @return Returns the builder.
         * @see #servicePredicate(BiPredicate)
         */
        public Builder operationPredicate(OperationPredicate operationPredicate) {
            this.operationPredicate = Objects.requireNonNull(operationPredicate);
            servicePredicate = (model, service) -> false;
            return this;
        }

        /**
         * Configures a predicate that makes a plugin only apply to a set of
         * operations that match one or more of the set of given shape names,
         * and ensures that the plugin is not applied globally to services.
         *
         * <p>By default, a plugin applies globally to a service, which thereby
         * applies to every operation when the middleware stack is copied.
         *
         * @param operationNames Set of operation names.
         * @return Returns the builder.
         */
        public Builder appliesOnlyToOperations(Set<String> operationNames) {
            operationPredicate((model, service, operation) -> operationNames.contains(operation.getId().getName()));
            return servicePredicate((model, service) -> false);
        }

        /**
         * Configures a predicate that applies the plugin to a service if the
         * predicate matches a given model and service.
         *
         * <p>When this method is called, the {@code operationPredicate} is
         * automatically configured to return false for every operation,
         * causing the plugin to only apply to services and not to individual
         * operations.
         *
         * <p>By default, a plugin applies globally to a service, which
         * thereby applies to every operation when the middleware stack is
         * copied. Setting a custom service predicate is useful for plugins
         * that should only be applied to specific services or only applied
         * at the operation level.
         *
         * @param servicePredicate Service predicate.
         * @return Returns the builder.
         */
        public Builder servicePredicate(BiPredicate<Model, ServiceShape> servicePredicate) {
            this.servicePredicate = Objects.requireNonNull(servicePredicate);
            operationPredicate = (model, service, operation) -> false;
            return this;
        }

        /**
         * Sets the config fields that will be added to the client config by this plugin.
         *
         * <p>Each config field will be added to the client's Config object and will
         * result in a corresponding getter method being added to the config. E.g.:
         * <p>
         * type ClientOptions struct {
         * // My docs.
         * MyField string
         * }
         * <p>
         * func (o ClientOptions) GetMyField() string {
         * return o.MyField
         * }
         *
         * @param configFields The config fields to add to the client config.
         * @return Returns the builder.
         */
        public Builder configFields(Collection<ConfigField> configFields) {
            this.configFields = new HashSet<>(configFields);
            return this;
        }

        /**
         * Adds a config field that will be added to the client config by this plugin.
         *
         * <p>Each config field will be added to the client's Config object and will
         * result in a corresponding getter method being added to the config. E.g.:
         * <p>
         * type ClientOptions struct {
         * // My docs.
         * MyField string
         * }
         * <p>
         * func (o ClientOptions) GetMyField() string {
         * return o.MyField
         * }
         *
         * @param configField The config field to add to the client config.
         * @return Returns the builder.
         */
        public Builder addConfigField(ConfigField configField) {
            this.configFields.add(configField);
            return this;
        }

        /**
         * Sets the config field resolvers that will be added to the client by this plugin.
         *
         * @param configFieldResolvers The config field resolvers.
         * @return Returns the builder.
         */
        public Builder configFieldResolvers(Collection<ConfigFieldResolver> configFieldResolvers) {
            this.configFieldResolvers = new HashSet<>(configFieldResolvers);
            return this;
        }

        /**
         * Adds a config field resolver that will be added to the client by this plugin.
         *
         * @param configFieldResolver The config field resolver.
         * @return Returns the builder.
         */
        public Builder addConfigFieldResolver(ConfigFieldResolver configFieldResolver) {
            this.configFieldResolvers.add(configFieldResolver);
            return this;
        }
    }
}