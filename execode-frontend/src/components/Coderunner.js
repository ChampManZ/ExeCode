import CoderunnerHeader from "./CoderunnerHeader"
import CoderunnerEditor from "./CoderunnerEditor"

const flexboxStyle = {
  display: "flex",
  justifyContent: "space-evenly"
}

function Coderunner() {
  return (
    <div>
        <CoderunnerHeader />
        <div style={flexboxStyle}>
          <div></div>
          <CoderunnerEditor />
          <div></div>
        </div>
    </div>
  )
}

export default Coderunner