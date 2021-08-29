import styles from './table.module.scss'
import React from 'react'

const EditableCell = ({
    value: initialValue,
    row: { index },
    column: { id },
  }: any) => {
    // We need to keep and update the state of the cell normally
    // const [value, setValue] = React.useState(initialValue)
  
    // // If the initialValue is changed external, sync it up with our state
    // React.useEffect(() => {
    //   setValue(initialValue)
    // }, [initialValue])

    // if (isEditingCell)
    //   return <input className="input is-small" type="text" value={value} onChange={onChange} />
    return(initialValue ? initialValue : <span className={styles.nullValue}>NULL</span>) 
}

export default EditableCell
  