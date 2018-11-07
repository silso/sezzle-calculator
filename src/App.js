class App extends React.Component {
	render() {
		return <Calculator/>;
	}
}

class MathLog extends React.Component {
	constructor(props) {
		super(props);
		this.props.mathLog = [];
	}

	render() {
		const reverseLog = this.props.mathLog.slice(0).reverse();
		const list = reverseLog.map((math, num) => {
			return <p>{math}</p>
		});
		console.log("math log: ", list);
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

		this.submit = this.submit.bind(this);
		let source = new EventSource("http://localhost:10000/stream");
		source.addEventListener('message', e => {
		  console.log("sse: ", e.data);
			let pushed = this.state.mathLog.concat(e.data);
			if (pushed.length > 10) {
				pushed.splice(0, 1);
			}
			this.setState({mathLog: pushed});
		}, false);
	}

	submit(e) {
		e.preventDefault();
		let form = new FormData(e.target);
		let mathMessage = form.get("mathInput");
		let arithmeticRegex = /^(\d+(\.\d+)?)([-+*/](\d+(\.\d+)?))*$/;
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
