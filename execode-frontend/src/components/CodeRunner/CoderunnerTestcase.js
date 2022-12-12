// import CoderunnerUploadEditor from "./CoderunnerUploadEditor"
import React, { useState } from 'react'
// import axios from 'axios'
import CodeMirror from '@uiw/react-codemirror'
import CoderunnerBtn from './CoderunnerBtn'

const paddingStyle = {
    padding: "4px 20px"
}

const uploadBtnStyle = {
    backgroundColor: "#F9A084",
    color: "white",
    padding: "0.5rem",
    borderRadius: "0.3rem",
    cursor: "pointer",
    marginTop: "1rem"
}

const divPaddingStyle = {
    padding: "20px 2px"
}

function CoderunnerTestcase(props) {

    const [isClick, setIsClick] = useState(false)
    const [stateTestCases, setStateTestCases] = useState([])
    const [myLength, setMyLength] = useState(0)
    let testCases = []
    const [code, setCode] = useState("// Write your test case function here\n")

    const onChange = React.useCallback((value, viewUpdate) => {
      setCode(value)
    }, []);
    
    // const testCaseURL = ""
    const submitTestCase = () => {
        testCases = code.split(";")

        setStateTestCases(testCases)
        setIsClick(false)
        setMyLength(testCases.length)
        props.sendTestCase(testCases)

        console.log(testCases)
    }

    const allTestCase = stateTestCases.map((element, key) => {
        return (
            <div key={key}>
                <p>{element}</p>
            </div>
        )
    })

    // props.sendTestCase(testCases)

    return (
        <div style={paddingStyle}>
            <h2>Test Case</h2>
            <p>To add test cases, please use ";" to separate between cases.</p>
            <p>Example: [2, 5];[5,-10]</p>

            {/* <CoderunnerUploadForm /> */}

            { isClick ? 
            <div style={divPaddingStyle}>
                <CodeMirror className='editor-test'
                height="auto"
                width='auto'
                value={code}
                onChange={onChange}
                />
                <CoderunnerBtn name="Add" style={uploadBtnStyle} clickFunc={submitTestCase} />
            </div> :
            <CoderunnerBtn name="Add Test Case" style={uploadBtnStyle} clickFunc={(event) => { setIsClick(true) }} />
            }
            {`Input n = ${myLength}`}
            {allTestCase}
        </div>
    )
}

export default CoderunnerTestcase