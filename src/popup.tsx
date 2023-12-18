import { render } from "preact"

import "./popup.css"

export const App = () => {
	return (
		<div className="container">
			<img src="/assets/icon128.png" />
			<h1>AdBrick</h1>
			<p>
				A 100% free ad blocker to protect you and enhance your browsing
				experience
			</p>
			<p>
				<a href="https://adbrick.org" target="_blank">
					adbrick.org
				</a>
			</p>
		</div>
	)
}

render(<App />, document.getElementById("app")!)
