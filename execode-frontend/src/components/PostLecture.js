import {useState} from 'react'
import axios from 'axios'
// Import Worker
import { Worker } from '@react-pdf-viewer/core';
// Import the main Viewer component
import { Viewer } from '@react-pdf-viewer/core';
// Import the styles
import '../styles/lecturer.css';
// default layout plugin
import { defaultLayoutPlugin } from '@react-pdf-viewer/default-layout';
// Import styles of default layout plugin
import '@react-pdf-viewer/default-layout/lib/styles/index.css';

function PostLecture() {

  // creating new plugin instance
  const defaultLayoutPluginInstance = defaultLayoutPlugin();

  // pdf file onChange state
  const [pdfFile, setPdfFile]=useState(null);

  // pdf file error state
  const [pdfError, setPdfError]=useState('');

  // handle file onChange event
  const allowedFiles = ['application/pdf'];
  const handleFile = (e) =>{
    let selectedFile = e.target.files[0];
    // console.log(selectedFile.type);
    if(selectedFile){
      if(selectedFile&&allowedFiles.includes(selectedFile.type)){
        let reader = new FileReader();
        reader.readAsDataURL(selectedFile);
        reader.onloadend=(e)=>{
          setPdfError('');
          setPdfFile(e.target.result);
        }
      }
      else{
        setPdfError('Not a valid pdf: Please select only PDF');
        setPdfFile('');
      }
    }
    else{
      console.log('please select a PDF');
    }
  }

  const handleFileSubmission = () => {
    if (pdfFile !== null) {
      let formdata = new FormData();
      formdata.append('lecture-file', pdfFile)
      formdata.append('className', 'newClass')
      formdata.append('fileName', 'sampleFile')
      formdata.append('module', 1)

      axios.post(
        "http://localhost:3000/pdflecture",
          formdata,
          {
              headers: {
                  "Content-Type": "multipart/form-data",
              },
          }
      )
      .then(res => {
          console.log(`Success` + res.data);
      })
      .catch(err => {
          console.log(err);
      })
    } else {
      console.log("missing file")
    }
  }

  return (
    <div className="container">

      {/* Upload PDF */}
      <form>

        <label><h5>Upload PDF</h5></label>
    
        <input type='file' className="form-control"
        onChange={handleFile}></input>

        {/* we will display error message in case user select some file
        other than pdf */}
        {pdfError&&<span className='text-danger'>{pdfError}</span>}

      </form>

      {/* View PDF */}

      <form>
            <h5>View PDF</h5>
            <div className="viewer">

                {/* render this if we have a pdf file */}
                {pdfFile&&(
                <Worker workerUrl="https://unpkg.com/pdfjs-dist@3.1.81/build/pdf.worker.min.js">
                    <Viewer fileUrl={pdfFile}
                    plugins={[defaultLayoutPluginInstance]}></Viewer>
                </Worker>
                )}

                {/* render this if we have pdfFile state null   */}
                {!pdfFile&&<>No file is selected yet</>}

            </div>
      </form>

      {/* Submit button */}
      <button type="button" className="lecture-button" onClick={handleFileSubmission}>Submit</button>

    </div>
  );
}

export default PostLecture;
