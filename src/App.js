class App extends React.Component {
	render() {
		return <Calculator/>;
	}
}


class MathLog extends React.Component {
	constructor(props) {
		super(props);
		//stores the history of calculations done by this client and others
		this.props.mathLog = [];
	}

	render() {
		//the mathLog is reversed only for rendering since it's easier to store it the way it is
		const reverseLog = this.props.mathLog.slice(0).reverse();
		const list = reverseLog.map((math, num) => {
			return <p>{math}</p>
		});
		return (
			<div>{list}</div>
		);
	}
}


class Calculator extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			mathLog: [],
		}
		this.submit = this.submit.bind(this);

		//the server responds to this post request with it's version of mathLog
		//this should only run once per client
		$.post(
			"/first-post/",
			{},
			data => {
				//remove empty math logs
				for (let i = 0; i < 10; i++) {
					if (data.mathLog[i] == "") {
						data.mathLog.splice(i)
						break;
					}
				}
				this.setState({mathLog: data.mathLog});
			},
			"json"
		);

		//this listens to Server-Sent Events from /stream and adds the data from those to mathLog
		let source = new EventSource("http://localhost:10000/stream");
		source.addEventListener("message", e => {
			let pushed = this.state.mathLog.concat(e.data);
			if (pushed.length > 10) {
				//when it goes past 10, start trimming the first elements of the array off
				pushed.splice(0, 1);
			}
			this.setState({mathLog: pushed});
		}, false);
	}

	submit(e) {
		e.preventDefault();
		let form = new FormData(e.target);
		let mathMessage = form.get("mathInput");
		//matches squences of decimal/integers (# or #.#) joined by +,-,*,/
		let arithmeticRegex = /^(\d+(\.\d+)?)([-+*/](\d+(\.\d+)?))*$/;
		//this error handling should probably be done on the server
		if (!mathMessage.match(arithmeticRegex)) {
			console.log("bad math");
			return;
		}
		//oooo spooky eval
		let answer = eval(mathMessage);
		mathMessage += " = " + answer;
		document.getElementById("mathInput").value = "";
		$.post(
			"/calculator-post/",
			{mathToServer: mathMessage},
		);
	}
	
  render() {
    return (
			<div>
				<h1>Calculator</h1>
				<form autocomplete="off" onSubmit={this.submit}>
					<input type="text" id="mathInput" name="mathInput" placeholder="feed me math"/>
					<input type="submit" value="="/>
				</form>
				<MathLog mathLog={this.state.mathLog}/>
			</div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById("app"));
