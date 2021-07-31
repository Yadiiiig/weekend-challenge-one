package router

import (
	"net/http"
)

func (c *Collection) GetHelp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
		<html>
			<head>
				<title>KMA -Help</title>
			</head>
			<body>
				<h1>HELP_ME_KMA</h1>
				<ul>
					<li><b>get_time_left</b> - To get the timer.</li>
					<li><b>get_challenge</b> - To get the current challenge</li>
					<li><b>submit_challenge</b> - To submit your answer</li>
					<li><b>credits</b> - Information about this project</li>
				</ul>
			</body>
		</html>
	`))
}
