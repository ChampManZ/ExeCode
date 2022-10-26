import CoderunnerHeader from "./CoderunnerHeader"
import CoderunnerEditor from "./CoderunnerEditor"
import CoderunnerOutput from "./CoderunnerOutput"
import CoderunnerTestcase from "./CoderunnerTestcase"

const flexboxStyle = {
  display: "flex",
  justifyContent: "space-evenly"
}

function Coderunner() {
  return (
    <div>
        <CoderunnerHeader />
        <div style={flexboxStyle}>
          <CoderunnerTestcase />
          <CoderunnerEditor />
          <CoderunnerOutput />
        </div>
    </div>
  )
}

export default Coderunner