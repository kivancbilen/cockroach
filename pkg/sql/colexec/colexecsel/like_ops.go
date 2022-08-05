// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecsel

import (
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/colexeccmp"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecop"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/eval"
	"github.com/cockroachdb/errors"
)

// GetLikeOperator returns a selection operator which applies the specified LIKE
// pattern, or NOT LIKE if the negate argument is true. The implementation
// varies depending on the complexity of the pattern.
func GetLikeOperator(
	ctx *eval.Context, input colexecop.Operator, colIdx int, pattern string, negate bool,
) (colexecop.Operator, error) {
	likeOpType, patterns, err := colexeccmp.GetLikeOperatorType(pattern)
	if err != nil {
		return nil, err
	}
	pat := patterns[0]
	base := selConstOpBase{
		OneInputHelper: colexecop.MakeOneInputHelper(input),
		colIdx:         colIdx,
	}
	switch likeOpType {
	case colexeccmp.LikeAlwaysMatch:
		// Use an empty prefix operator to get correct NULL behavior.
		return &selPrefixBytesBytesConstOp{
			selConstOpBase: base,
			constArg:       []byte{},
			negate:         negate,
		}, nil
	case colexeccmp.LikeConstant:
		if negate {
			return &selNEBytesBytesConstOp{
				selConstOpBase: base,
				constArg:       pat,
			}, nil
		}
		return &selEQBytesBytesConstOp{
			selConstOpBase: base,
			constArg:       pat,
		}, nil
	case colexeccmp.LikeContains:
		return &selContainsBytesBytesConstOp{
			selConstOpBase: base,
			constArg:       pat,
			negate:         negate,
		}, nil
	case colexeccmp.LikePrefix:
		return &selPrefixBytesBytesConstOp{
			selConstOpBase: base,
			constArg:       pat,
			negate:         negate,
		}, nil
	case colexeccmp.LikeRegexp:
		re, err := eval.ConvertLikeToRegexp(ctx, string(patterns[0]), false, '\\')
		if err != nil {
			return nil, err
		}
		return &selRegexpBytesBytesConstOp{
			selConstOpBase: base,
			constArg:       re,
			negate:         negate,
		}, nil
	case colexeccmp.LikeSkeleton:
		return &selSkeletonBytesBytesConstOp{
			selConstOpBase: base,
			constArg:       patterns,
			negate:         negate,
		}, nil
	case colexeccmp.LikeSuffix:
		return &selSuffixBytesBytesConstOp{
			selConstOpBase: base,
			constArg:       pat,
			negate:         negate,
		}, nil
	default:
		return nil, errors.AssertionFailedf("unsupported like op type %d", likeOpType)
	}
}
