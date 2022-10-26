import { useState } from "react"
import CoderunnerUploadForm from "./CoderunnerUploadForm"

function CoderunnerTestcase() {

    const [testCase, setTestCase] = useState({})

    return (
        <div>
            <h2>Test Case</h2>
            <CoderunnerUploadForm />
        </div>
    )
}

export default CoderunnerTestcase