import styles from './jsontable.module.scss'
import React, { useRef } from 'react'
import { js_beautify } from 'js-beautify'
import _ from 'lodash'
import toast from 'react-hot-toast'
import ReactJson from 'react-json-view'
import ReactCodeMirror, { ReactCodeMirrorRef } from '@uiw/react-codemirror'
import { javascript } from '@codemirror/lang-javascript'
import Button from '../../ui/Button'


const JsonCell = ({
    row: { index, original },
    editingCellIndex,
    startEditing,
    onSaveCell
}: any) => {
    // We need to keep and update the state of the cell normally
    const [value, setValue] = React.useState(original)
    const [editingValue, setEditingValue] = React.useState<string>(JSON.stringify(original))

    const editorRef = useRef<ReactCodeMirrorRef | null>(null);

    // If the initialValue is changed external, sync it up with our state
    React.useEffect(() => {
        setValue(original)
        setEditingValue(JSON.stringify(original))
    }, [original])

    const onChange = React.useCallback((value: any) => {
        setEditingValue(value)
    }, []);

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
            <ReactCodeMirror
                ref={editorRef}
                value={js_beautify(editingValue)}
                extensions={[javascript()]}
                onChange={onChange}
            />
            <div className="column is-flex is-justify-content-flex-end">
                <Button
                    className="is-small" 
                    icon={<i className="fas fa-window-close"/>} 
                    onClick={cancelEdit}
                />
                <span>&nbsp;&nbsp;</span>
                <Button 
                    className="is-primary is-small" 
                    icon={<i className="fas fa-save"/>} 
                    onClick={onSave}
                />
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