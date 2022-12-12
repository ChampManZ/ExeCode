import CoderunnerHeader from "../components/CodeRunner/CoderunnerHeader"
import CoderunnerEditor from "../components/CodeRunner/CoderunnerEditor"
import CoderunnerOutput from "../components/CodeRunner/CoderunnerOutput"
import CoderunnerTestcase from "../components/CodeRunner/CoderunnerTestcase"
import Sidebar from "../components/Sidebar"
import Split from "react-split"
import "../styles/splitter.css"
import { Box } from "@mui/system"
import { useState } from "react"

const flexboxStyle = {
  display: "flex",
  justifyContent: "space-evenly"
}

const bottomBorderStyle = {
  borderBottom: "1px solid #E5E5E5"
}

function Coderunner(props) {
  const [data, setData] = useState('')
  const [output, setOutput] = useState('')
  const [testCases, setTestCase] = useState([])

  const sendData = (data) => {
    setData(data)
  }

  const outputData = (output) => {
    setOutput(output)
  }

  const sendTestCase = (testCases) => {
    setTestCase(testCases)
  }

  return (
    <>
    
    <Sidebar></Sidebar>
    <Box sx={{ mt:15 }}>
      <CoderunnerHeader sendData={sendData} />
      <Split sizes={[20, 60, 20]} cursor="col-resize" direction="horizontal" style={flexboxStyle}>
          <CoderunnerTestcase sendTestCase={sendTestCase} />
          <CoderunnerEditor data={data} outputData={outputData} testCases={testCases} />
          <CoderunnerOutput output={output} />
      </Split>
      <div style={bottomBorderStyle} />
    </Box>
    </>
    
  )
}

export default Coderunner