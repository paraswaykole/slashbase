import Button from '../../ui/Button'
import styles from './table.module.scss'
import React from 'react'

const EditableCell = ({
  value: initialValue,
  row: { index, original },
  column: { id },
  editCell,
  resetEditCell,
  onSaveCell
}: any) => {

  initialValue = Array.isArray(initialValue) ?
    `{${initialValue.join(",")}}` : initialValue

  initialValue = (initialValue !== null && typeof initialValue === "object") ?
    JSON.stringify(initialValue) : initialValue

  // We need to keep and update the state of the cell normally
  const [value, setValue] = React.useState(initialValue)

  // If the initialValue is changed external, sync it up with our state
  React.useEffect(() => {
    setValue(initialValue)
  }, [initialValue])

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setValue(e.target.value)
  }

  const cancelEdit = () => {
    setValue(initialValue)
    resetEditCell()
  }

  const onSave = async () => {
    onSaveCell(index, original, id, value)
  }

  const isEditingCell = editCell.length == 2 && editCell[0] === index && editCell[1] === id

  if (isEditingCell) {
    return (
      <div className="field has-addons">
        <div className="control is-expanded">
          <input
            className={"input is-small " + styles.cellinput}
            type="text"
            placeholder={"Enter " + id}
            value={value}
            onChange={onChange} />
        </div>
        <div className="control">
          <Button className='is-small' icon={<i className="fas fa-check"/>} onClick={onSave}/>
        </div>
        <div className="control">
          <Button className='is-small' icon={<i className="fas fa-times"/>} onClick={cancelEdit}/>
        </div>
      </div>
    )
  }

  return (initialValue === null ?
    <span className={styles.nullValue}>NULL</span> :
    initialValue === false ?
      <span>false</span> :
      initialValue === true ?
        <span>true</span> :
        initialValue
  )
}

export default EditableCell
