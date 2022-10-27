import CoderunnerHeader from "./CoderunnerHeader"
import CoderunnerEditor from "./CoderunnerEditor"
import CoderunnerOutput from "./CoderunnerOutput"
import CoderunnerTestcase from "./CoderunnerTestcase"
import Split from "react-split"

const flexboxStyle = {
  display: "flex",
  justifyContent: "space-evenly"
}

function Coderunner() {
  return (
    <div>
        <CoderunnerHeader />
        <Split sizes={[30, 50, 30]} direction="vertical" style={flexboxStyle} minSize={30} >
            <CoderunnerTestcase />
            <CoderunnerEditor />
            <CoderunnerOutput />
        </Split>
    </div>
  )
}

export default Coderunner