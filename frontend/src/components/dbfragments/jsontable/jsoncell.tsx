import styles from './jsontable.module.scss'
import React, { useRef } from 'react'
import dynamic from 'next/dynamic'
import ReactJson from 'react-json-view'
import jsonBeautify from 'json-beautify'
import _ from 'lodash'
import toast from 'react-hot-toast'

// temp hack
declare module 'json-beautify' {
    export default function beautify(value: any, replacer: null, space: number | string, limit?: number): string
}

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

const JsonCell = ({
    row: { index, original },
    editingCellIndex,
    startEditing,
    onSaveCell
}: any) => {
    // We need to keep and update the state of the cell normally
    const [value, setValue] = React.useState(original)
    const [editingValue, setEditingValue] = React.useState<string>(JSON.stringify(original))

    const editorRef = useRef<ReactCodeMirror.ReactCodeMirror | null>(null);

    // If the initialValue is changed external, sync it up with our state
    React.useEffect(() => {
        setValue(original)
    }, [original])

    const isEditing: boolean = editingCellIndex == index

    if (isEditing) {

        const cancelEdit = () => {
            setValue(original)
            startEditing(null)
        }

        const onSave = async () => {
            let jsonData: any
            try {
                jsonData = JSON.parse(editingValue)
            } catch (e: any) {
                toast.error(e.message)
                return
            }
            onSaveCell(jsonData._id, JSON.stringify(_.omit(jsonData, ['_id'])))
        }

        return (<React.Fragment>
            <ForwardRefCodeMirror
                ref={editorRef}
                value={jsonBeautify(value, null, 2)}
                options={{
                    mode: 'javascript',
                    theme: 'duotone-light',
                    lineNumbers: true
                }}
                onChange={(newValue) => {
                    setEditingValue(newValue)
                }} />
            <div className="column is-flex is-justify-content-flex-end">
                <button className="button is-small" onClick={cancelEdit}>
                    <span className="icon is-small">
                        <i className="fas fa-window-close" />
                    </span>
                </button>
                <span>&nbsp;&nbsp;</span>
                <button className="button is-primary is-small" onClick={onSave}>
                    <span className="icon is-small">
                        <i className="fas fa-save" />
                    </span>
                </button>
            </div>
        </React.Fragment>)
    }

    return (<React.Fragment>
        <ReactJson
            src={value}
            name={null}
            collapsed={1}
            displayDataTypes={false}
            displayObjectSize={false}
            enableClipboard={false}
        />
    </React.Fragment>)
}

export default JsonCell