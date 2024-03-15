package main

var templateToUse = `
	<h1>Issues<h1>
	<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>
	{{range .issues}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	<h1>Milestones<h1>
	<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>Creator</th>
		<th>Title</th>
	</tr>
	{{range .milestones}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.Creator.HTMLURL}}'>{{.Creator.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	<h1>Users<h1>
	<table>
	<tr style='text-align: left'>
		<th>User</th>
	</tr>
	{{range .users}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
	<tr>
	{{end}}
	</table>
`
