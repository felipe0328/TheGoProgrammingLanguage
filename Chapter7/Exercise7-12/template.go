package main

var templateToUse = `
	<h1> Elements </h1>
	<table>
		<tr>
			<th>Item</th>
			<th>Value</th>
		</tr>
		{{range $key, $value := .}}
		<tr>
			<td>{{$key}}</td>
			<td>{{$value}}</td>
		</tr>
		{{end}}
	</table>
`
