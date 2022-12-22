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
    onSaveCell(original["0"], id, value)
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
          <button className="button is-small" onClick={onSave}>
            <span className="icon is-small">
              <i className="fas fa-check"></i>
            </span>
          </button>
        </div>
        <div className="control">
          <button className="button is-small" onClick={cancelEdit}>
            <span className="icon is-small">
              <i className="fas fa-times"></i>
            </span>
          </button>
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
