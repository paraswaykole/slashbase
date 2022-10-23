import styles from './table.module.scss'
import React from 'react'
import ReactJson from 'react-json-view'

const JsonCell = ({
    row: { original }
}: any) => {
    // We need to keep and update the state of the cell normally
    const [value, setValue] = React.useState(original)

    // If the initialValue is changed external, sync it up with our state
    React.useEffect(() => {
        setValue(original)
    }, [original])

    return (<ReactJson
        src={value}
        name={null}
        collapsed={1}
        displayDataTypes={false}
        displayObjectSize={false}
        enableClipboard={false}
    />)
}

export default JsonCell
