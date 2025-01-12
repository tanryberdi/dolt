// Copyright 2022 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jwtauth

import (
	"time"

	"gopkg.in/square/go-jose.v2/jwt"
)

type JWTValidator interface {
	ValidateJWT(unparsed string, reqTime time.Time) (*Claims, error)
}

type fetchingJWTValidator struct {
	jwks     *fetchedJWKS
	expected jwt.Expected
}

func NewJWTValidator(provider JWTProvider) JWTValidator {
	expected := jwt.Expected{Issuer: provider.Issuer, Audience: jwt.Audience{provider.Audience}}
	return &fetchingJWTValidator{jwks: newJWKS(provider), expected: expected}
}

func (v *fetchingJWTValidator) ValidateJWT(unparsed string, reqTime time.Time) (*Claims, error) {
	return ValidateJWT(unparsed, reqTime, v.jwks, v.expected)
}
