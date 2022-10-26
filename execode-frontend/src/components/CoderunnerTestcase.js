import { useState } from "react"
import CoderunnerUploadForm from "./CoderunnerUploadForm"

function CoderunnerTestcase() {

    const [testCase, setTestCase] = useState({})

    return (
        <div>
            <h3>Test Case</h3>
            <CoderunnerUploadForm />
        </div>
    )
}

export default CoderunnerTestcase