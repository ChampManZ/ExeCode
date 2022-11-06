import CoderunnerHeader from "../components/CodeRunner/CoderunnerHeader"
import CoderunnerEditor from "../components/CodeRunner/CoderunnerEditor"
import CoderunnerOutput from "../components/CodeRunner/CoderunnerOutput"
import CoderunnerTestcase from "../components/CodeRunner/CoderunnerTestcase"
import Split from "react-split"
import "../styles/splitter.css"

const flexboxStyle = {
  display: "flex",
  justifyContent: "space-evenly"
}

const bottomBorderStyle = {
  borderBottom: "1px solid #E5E5E5"
}

function Coderunner() {
  return (
    <div>
        <CoderunnerHeader />
        <Split sizes={[10, 70, 20]} cursor="col-resize" direction="horizontal" style={flexboxStyle}>
            <CoderunnerTestcase />
            <CoderunnerEditor />
            <CoderunnerOutput />
        </Split>
        <div style={bottomBorderStyle} />
    </div>
  )
}

export default Coderunner