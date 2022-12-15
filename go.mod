module github.com/PauloPortugal/gogen

go 1.19

// [CVE-2020-14040] CWE-835: Loop with Unreachable Exit Condition ('Infinite Loop')
// [CVE-2022-32149] CWE-772: Missing Release of Resource after Effective Lifetime
// [CVE-2021-38561] CWE-125: Out-of-bounds Read
replace golang.org/x/text => golang.org/x/text v0.5.0

require github.com/smartystreets/goconvey v1.7.2

require (
	github.com/gopherjs/gopherjs v0.0.0-20181017120253-0766667cb4d1 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/smartystreets/assertions v1.2.0 // indirect
)
