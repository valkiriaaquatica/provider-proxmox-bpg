// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

// SchemeBuilder collects the scheme builder functions for this API group.
var SchemeBuilder = runtime.NewSchemeBuilder()

// AddToScheme applies all additions to the provided scheme.
var AddToScheme = SchemeBuilder.AddToScheme
