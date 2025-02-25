// Copyright (c) 2019, Sylabs Inc. All rights reserved.
// Copyright (c) 2020-2025, NVIDIA CORPORATION. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package hpcjob

type Status struct {
	Code int
	Str  string
}

const (
	JOB_STATUS_UNKNOWN = iota
	JOB_STATUS_PENDING
	JOB_STATUS_QUEUED
	JOB_STATUS_RUNNING
	JOB_STATUS_STOP
	JOB_STATUS_DONE
)

var StatusUnknown = Status{
	Code: JOB_STATUS_UNKNOWN,
	Str:  "UNKNOWN",
}
var StatusPending = Status{
	Code: JOB_STATUS_PENDING,
	Str:  "PENDING",
}
var StatusQueued = Status{
	Code: JOB_STATUS_QUEUED,
	Str:  "QUEUED",
}
var StatusRunning = Status{
	Code: JOB_STATUS_RUNNING,
	Str:  "RUNNING",
}
var StatusStop = Status{
	Code: JOB_STATUS_STOP,
	Str:  "STOPPED",
}
var StatusDone = Status{
	Code: JOB_STATUS_DONE,
	Str:  "DONE",
}
