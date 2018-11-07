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
		const list = this.props.mathLog.map((math, num) => {
			return <p>{`${num}: ${math}`}</p>
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
			newMathLog: []
			reading: false,
		}
		this.submit = this.submit.bind(this);
		let source = new EventSource("http://localhost:10000/stream");
		source.addEventListener('message', function(e) {
		  console.log("sse: ", e.data);
			if (e.data === "START") {
				this.setState({reading: true, newMathLog: []})
			}
			else if (e.data === "IGNORE") {
				this.setState({reading: false})
			}
			else if (this.state.reading) {
				this.state.newMathLog.push(e.data)
			}
		}, false);
	}

	submit(e) {
		e.preventDefault();
		let mathMessage = (new FormData(e.target)).get("mathInput");
		$.post(
			"/calculator-post/",
			{mathToServer: mathMessage},
			/*(function(data) {
				this.state.mathLog.push(data.mathToClients);
				this.setState({mathLog: this.state.mathLog});
			}).bind(this),
			"json"*/
		);
	}
	
  render() {
    return (
			<div>
				<h1>Calculator</h1>
				<form autocomplete="off" onSubmit={this.submit}>
					<input type="text" name="mathInput" placeholder="feed me math"/>
					<input type="submit" value="="/>
				</form>
				<MathLog mathLog={this.state.mathLog}/>
			</div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById("app"));
