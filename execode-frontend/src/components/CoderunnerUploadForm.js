import { useState } from "react"

function CoderunnerUploadForm() {
    // Didn't provide action path to backend yet. Don't forget to do it.
    // Didn't do handle submission after POST to backend yet too.
    const [isUpload, setIsUpload] = useState(false)
    const [fileDetail, setFileDetail] = useState()

    const fileHandler = (event) => {
        setFileDetail(event.target.files[0])
        setIsUpload(true)
    }

    const uploadBtnStyle = {
        backgroundColor: "#F9A084",
        color: "white",
        padding: "0.5rem",
        borderRadius: "0.3rem",
        cursor: "pointer",
        marginTop: "1rem"
    }

    // Don't forget to add this
    const fetchURL = ""
    const handleSubmit = async (event) => {
        event.preventDefault();
        const formData = new FormData()
        formData.append('title', fileDetail.name)
        formData.append('filepath', fileDetail)
        try {
            let res = await fetch(fetchURL, {
                method: "POST",
                body: formData
            })
            console.log(res.json())
        } catch(err) {
            console.log(err)
        }
    }

    return (
        <div>
            <form onSubmit={handleSubmit}>
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