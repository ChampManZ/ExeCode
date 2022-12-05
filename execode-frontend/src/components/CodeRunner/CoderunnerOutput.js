const paddingStyle = {
  padding: "4px 20px"
}

function CoderunnerOutput(props) {

  return (
    <div style={paddingStyle}>
        <h2>Output</h2>
        <p id="output-id"></p>
        <p id="error"></p>
    </div>
  )
}

export default CoderunnerOutput