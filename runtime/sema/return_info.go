/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sema

import (
	"github.com/onflow/cadence/runtime/common/persistent"
)

// TODO: rename to e.g. ControlFlowInfo

// ReturnInfo tracks control-flow information
type ReturnInfo struct {
	// MaybeReturned indicates that (the branch of) the function
	// contains a potentially taken return statement
	MaybeReturned bool
	// JumpOffsets contains the offsets of all jumps
	// (break or continue statements), potential or definite.
	//
	// If non-empty, indicates that (the branch of) the function
	// contains a potential break or continue statement
	JumpOffsets *persistent.OrderedSet[int]
	// DefinitelyReturned indicates that (the branch of) the function
	// contains a definite return statement
	DefinitelyReturned bool
	// DefinitelyHalted indicates that (the branch of) the function
	// contains a definite halt (a function call with a Never return type)
	DefinitelyHalted bool
	// DefinitelyJumped indicates that (the branch of) the function
	// contains a definite break or continue statement
	DefinitelyJumped bool
}

func NewReturnInfo() *ReturnInfo {
	return &ReturnInfo{
		JumpOffsets: persistent.NewOrderedSet[int](nil),
	}
}

func (ri *ReturnInfo) MaybeJumped() bool {
	return ri.JumpOffsets != nil &&
		!ri.JumpOffsets.IsEmpty()
}

func (ri *ReturnInfo) MergeBranches(thenReturnInfo *ReturnInfo, elseReturnInfo *ReturnInfo) {
	ri.MaybeReturned = ri.MaybeReturned ||
		thenReturnInfo.MaybeReturned ||
		elseReturnInfo.MaybeReturned

	ri.DefinitelyReturned = ri.DefinitelyReturned ||
		(thenReturnInfo.DefinitelyReturned &&
			elseReturnInfo.DefinitelyReturned)

	ri.DefinitelyJumped = ri.DefinitelyJumped ||
		(thenReturnInfo.DefinitelyJumped &&
			elseReturnInfo.DefinitelyJumped)

	ri.DefinitelyHalted = ri.DefinitelyHalted ||
		(thenReturnInfo.DefinitelyHalted &&
			elseReturnInfo.DefinitelyHalted)
}

func (ri *ReturnInfo) MergePotentiallyUnevaluated(temporaryReturnInfo *ReturnInfo) {
	ri.MaybeReturned = ri.MaybeReturned ||
		temporaryReturnInfo.MaybeReturned

	// NOTE: the definitive return state does not change
}

func (ri *ReturnInfo) Clone() *ReturnInfo {
	result := NewReturnInfo()
	*result = *ri
	return result
}

func (ri *ReturnInfo) IsUnreachable() bool {
	return ri.DefinitelyReturned ||
		ri.DefinitelyHalted ||
		ri.DefinitelyJumped
}

func (ri *ReturnInfo) AddJumpOffset(offset int) {
	ri.JumpOffsets.Add(offset)
}

func (ri *ReturnInfo) WithNewJumpTarget(f func()) {
	ri.JumpOffsets = ri.JumpOffsets.Clone()
	f()
	ri.JumpOffsets = ri.JumpOffsets.Parent
}
