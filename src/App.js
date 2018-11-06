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
		}
		this.submit = this.submit.bind(this);
	}

	submit(e) {
		e.preventDefault();
		let mathMessage = (new FormData(e.target)).get("mathInput");
		$.post(
			"/calculator-post/",
			{mathToServer: mathMessage},
			(function(data) {
				this.state.mathLog.push(data.mathToClients);
				this.setState({mathLog: this.state.mathLog});
			}).bind(this),
			"json"
		);
	}
	
  render() {
    return (
			<div>
				<h1>Calculater</h1>
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
