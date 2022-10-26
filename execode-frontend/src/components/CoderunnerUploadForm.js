import { useState } from "react"

function CoderunnerUploadForm() {
    // Didn't provide action path to backend yet. Don't forget to do it.
    const [isUpload, setIsUpload] = useState(false)
    const [fileDetail, setFileDetail] = useState()

    const fileHandler = (event) => {
        setFileDetail(event.target.files[0])
        setIsUpload(true)
    }

    const handleSubmission = () => {

    }

    const uploadBtnStyle = {
        backgroundColor: "#F9A084",
        color: "white",
        padding: "0.5rem",
        borderRadius: "0.3rem",
        cursor: "pointer",
        marginTop: "1rem"
    }

    return (
        <div>
            <form>
                { isUpload ? 
                <div>
                    <h5>Sure to upload this test case?</h5>
                    <p>Name: { fileDetail.name }</p>
                    <p>Type: { fileDetail.type }</p>
                    <p>Size: { fileDetail.size } bytes</p>
                    <input type="submit" style={uploadBtnStyle} />
                </div> : 
                <div>
                    <input type="file" id="codeFileUpload" name="codeFileName" onChange={fileHandler} hidden />
                    <label htmlFor="codeFileUpload" style={uploadBtnStyle}>Upload via file</label>
                </div> 
                }
            </form>
        </div>
    )
}

export default CoderunnerUploadForm