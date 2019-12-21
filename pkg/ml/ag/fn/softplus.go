// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import (
	"brillion.io/spago/pkg/mat"
)

// SoftPlus(x) = 1/ β ​∗ log( 1 + exp(β ∗ x))
type SoftPlus struct {
	x         Operand
	beta      Operand
	threshold Operand
}

func NewSoftPlus(x, beta, threshold Operand) *SoftPlus {
	return &SoftPlus{x: x, beta: beta, threshold: threshold}
}

// Forward computes the output of the function.
func (r *SoftPlus) Forward() mat.Matrix {
	y := r.x.Value().ZerosLike()
	y.ApplyWithAlpha(softPlus, r.x.Value(), r.beta.Value().Scalar(), r.threshold.Value().Scalar())
	return y
}

func (r *SoftPlus) Backward(gy mat.Matrix) {
	if r.x.RequiresGrad() {
		gx := r.x.Value().ZerosLike()
		gx.ApplyWithAlpha(softPlusDeriv, r.x.Value(), r.beta.Value().Scalar(), r.threshold.Value().Scalar())
		gx.ProdInPlace(gy)
		r.x.PropagateGrad(gx)
	}
}