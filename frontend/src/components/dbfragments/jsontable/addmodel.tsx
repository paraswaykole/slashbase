import styles from './jsontable.module.scss'
import React, { useRef, useState } from 'react'
import dynamic from 'next/dynamic'
import { ApiResult, InsertedIDResponse, DBConnection, DBQueryData } from '../../../data/models'
import apiService from '../../../network/apiService'
import toast from 'react-hot-toast'

const WrappedCodeMirror = dynamic(() => {
    // @ts-ignore
    import('codemirror/mode/javascript/javascript')
    return import('../../lib/wrappedcodemirror')
}, { ssr: false })

const ForwardRefCodeMirror = React.forwardRef<
    ReactCodeMirror.ReactCodeMirror,
    ReactCodeMirror.ReactCodeMirrorProps
>((props, ref) => {
    return <WrappedCodeMirror {...props} editorRef={ref} />;
});

ForwardRefCodeMirror.displayName = 'ForwardRefCodeMirror';

type AddModal = {
    dbConnection: DBConnection
    mName: string,
    onAddData: (newData: any) => void
    onClose: () => void
}

const AddModal = ({ dbConnection, mName, onAddData, onClose }: AddModal) => {

    const editorRef = useRef<ReactCodeMirror.ReactCodeMirror | null>(null);
    const [newData, setNewData] = useState<any>(`{\n\t\n}`)

    const startAdding = async () => {
        let jsonData: any
        try {
            jsonData = JSON.parse(newData)
        } catch (e: any) {
            toast.error(e.message)
            return
        }
        const result: ApiResult<InsertedIDResponse> = await apiService.addDBData(dbConnection.id, "", mName, jsonData)
        if (result.success) {
            toast.success('data added')
            let mNewData = { _id: result.data.insertedId, ...jsonData }
            onAddData(mNewData)
            onClose()
        } else {
            toast.error(result.error!)
        }
    }

    return (<div className="modal is-active">
        <div className="modal-background"></div>
        <div className="modal-card">
            <header className="modal-card-head">
                <p className="modal-card-title">Add new data to {mName}</p>
                <button className="delete" aria-label="close" onClick={onClose}></button>
            </header>
            <section className="modal-card-body">
                <ForwardRefCodeMirror
                    ref={editorRef}
                    value={newData}
                    options={{
                        mode: 'javascript',
                        theme: 'duotone-light',
                        lineNumbers: true
                    }}
                    onChange={(newValue) => {
                        setNewData(newValue)
                    }} />
            </section>
            <footer className="modal-card-foot">
                <button className="button is-primary" onClick={startAdding}>Add</button>
                <button className="button" onClick={onClose}>Cancel</button>
            </footer>
        </div>
    </div>)
}

export default AddModal