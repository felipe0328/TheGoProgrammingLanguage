package main

var templateToUse = `
	<h1> Tracks </h1>
	<table>
		<tr>
			<th><a href="/?filter=Title">Title</a></th>
			<th>Artist</th>
			<th>Album</th>
			<th><a href="/?filter=Year">Year</a></th>
			<th>Length</th>
		</tr>
		{{range .}}
		<tr>
			<td>{{.Title}}</td>
			<td>{{.Artist}}</td>
			<td>{{.Album}}</td>
			<td>{{.Year}}</td>
			<td>{{.Length}}</td>
		</tr>
		{{end}}
	</table>
`
