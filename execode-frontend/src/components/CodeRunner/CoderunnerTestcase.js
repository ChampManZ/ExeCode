import CoderunnerUploadForm from "./CoderunnerUploadForm"
import CoderunnerUploadEditor from "./CoderunnerUploadEditor"

const paddingStyle = {
    padding: "4px 20px"
}

function CoderunnerTestcase() {
    return (
        <div style={paddingStyle}>
            <h2>Test Case</h2>
            <CoderunnerUploadForm />
            <CoderunnerUploadEditor />
        </div>
    )
}

export default CoderunnerTestcase