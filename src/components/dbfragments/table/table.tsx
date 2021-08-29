import styles from './table.module.scss'
import React, { useState } from 'react'
import { Cell, useTable } from 'react-table'
import toast from 'react-hot-toast';
import { DBConnection, DBQueryData } from '../../../data/models'
import EditableCell from './editablecell'
import apiService from '../../../network/apiService'


type ProjectCardPropType = { 
    queryData: DBQueryData,
    dbConnection: DBConnection
    mSchema: string,
    mName: string,
    isEditable: boolean,
    updateCellData: (oldCtid: string, newCtid: string, columnName: string, newValue: string|null|boolean)=>void
}

const Table = ({queryData, updateCellData, dbConnection, mSchema, mName, isEditable}: ProjectCardPropType) => {

    const [editCell, setEditCell] = useState<(string|number)[]>([])

    const data = React.useMemo(
        () => queryData.rows,
        [queryData]
    )

    const columns = React.useMemo(
        () => queryData.columns.filter(col => col !== 'ctid').map((col) => ({
            Header: col,
            accessor: col, 
        })),
        [queryData]
    )

    const defaultColumn = {
        Cell: EditableCell,
    }      

    const resetEditCell = ()=> {
        setEditCell([])
    }

    const onSaveCell = async (ctid: string, columnName: string, newValue: string) => {
        const result = await apiService.updateDBSingleData(dbConnection.id, mSchema, mName, ctid, columnName, newValue)
        if (result.success) {
            updateCellData(ctid, result.data.ctid, columnName, newValue)
            resetEditCell()
            toast.success('1 row updated');
        } else {
            toast.error(result.error!);
        }
    }

    const {
        getTableProps,
        getTableBodyProps,
        headerGroups,
        rows,
        prepareRow,
      } = useTable({ columns, data, defaultColumn, ...{ editCell, resetEditCell, onSaveCell }})

      const startEditing = (cell: Cell<any, any>)=>{
        if (isEditable)
            setEditCell([cell.row.index, cell.column.id])
    }
   
    return (
        <React.Fragment>
            <table {...getTableProps()} className={"table is-bordered is-striped is-narrow is-hoverable is-fullwidth"}>
                <thead>
                    {headerGroups.map(headerGroup => (
                        <tr {...headerGroup.getHeaderGroupProps()} key={headerGroup.id}>
                            {headerGroup.headers.map(column => (
                                <th {...column.getHeaderProps()} key={column.id}>
                                    {column.render('Header')}
                                </th>
                            ))}    
                        </tr>
                    ))}
                </thead>
                <tbody {...getTableBodyProps()}> 
                    {rows.map(row => {
                        prepareRow(row)
                        return (
                            <tr {...row.getRowProps()} key={row.id}>
                                {row.cells.map(cell => {
                                    return (<td {...cell.getCellProps()} onDoubleClick={()=>startEditing(cell)} key={row.id+""+cell.column.id}>
                                        {cell.render('Cell')}
                                    </td>
                                    )
                                })}
                            </tr>
                        )
                    })}
                </tbody>
            </table>
        </React.Fragment>
    )
}


export default Table