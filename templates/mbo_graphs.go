// Copyright 2017 Joyent, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package templates

const MboGraphsIndex = `
<html>
	<body>
		<h1>Conch : MBO Hardware Failures</h1>
		<h2>Text Reports</h2>

		<h3>Full Report</h3>
		<ul>
			<li><a href="/full">Text</a></li>
			<li><a href="/full.csv">CSV</a></li>
		</ul>

		<h3>Remediation Times</h3>
		<ul>
		{{ range .AzNames }}
			<li><a href="/reports/times/{{.}}">{{.}}</a></li>
		{{ end }}
		</ul>


		<h2>Graphs</h2>
		<h3>By Type</h3>
		<ul>
		{{ range .AzNames }}
			<li><a href="/graphics/{{.}}/by_type.png">{{.}}</a></li>
		{{ end }}
		</ul>

		<h3>By Vendor</h3>
		<ul>
		{{ range .AzNames }}
			<li><a href="/graphics/{{.}}/by_vendor.png">{{.}}</a></li>
		{{ end }}
		</ul>
	</body>
</html>
`

const MboGraphsReportsIndex = `
<html>
	<body>
		<h1>Conch: Hardware Failures for {{ .Name }}</h1>

		<ul>
		{{ range $type, $data := .Data.TimesByType }}
			<li><a href="/reports/times/{{ $.Name }}/{{ $type }}">{{ $type }}</a><ul>
				<li>Failure Count: {{$data.Count}}</li>
				<li>Mean: {{ $data.Mean }}</li>
				<li>Median: {{ $data.Median }}</li>
			</ul></li>
		{{ end }}
		</ul>

	</body>
</html>
`

const MboGraphsReportsBySubtype = `
<html>
	<body>
		<h1>Conch: Hardware Failures for {{.Az}}, Type {{.Name}} </h1>

		<ul>
		{{ range $type, $data := .Data }}
			<li><a href="/reports/times/{{ $.Az }}/{{ $.Name }}/{{ $type }}">{{ $type }}</a><ul>
				<li>Failure Count: {{$data.Count}}</li>
				<li>Mean: {{ $data.Mean }}</li>
				<li>Median: {{ $data.Median }}</li>
			</ul></li>
		{{ end }}
		</ul>

	</body>
</html>
`

const MboGraphsReportsByComponentAndSubtype = `
<html>
	<body>
		<h1>Conch: Hardware Failures for {{.Az}}, Type {{.Component}}, Subtype {{.Subtype}} </h1>

		<ul>
			<li>Failure Count: {{ .Data.Count }}</li>
			<li>Mean: {{ .Data.Mean }}</li>
			<li>Median: {{ .Data.Median }}</li>
		</ul>

		<h2>Affected Devices</h2>
		<ul>
		{{ range .Data.Devices }}
			<li><a href="https://conch.joyent.us/#!/device/{{.DeviceId}}" target="_blank">{{ .DeviceId }}</a><ul>
				<li>Remediation Time: {{ .RemediationTime }}</li>
				<li>Results:<ul>
					<li>First Failure:<ul>
						<li>Reported: {{ .FirstFail.Created }}</li>
						<li>Log: {{ .FirstFail.Result.Log }}</li>
					</ul></li>
					<li>First Pass:<ul>
						<li>Reported: {{ .FirstPass.Created }}</li>
						<li>Log: {{ .FirstPass.Result.Log }}</li>
					</ul></li>
				</ul></li>
			</ul></li>
		{{ end }}
		</ul>

	</body>
</html>

`