import styles from './table.module.scss'
import React from 'react'
import { useTable } from 'react-table'
import { DBQueryData } from '../../../data/models'
import EditableCell from './editablecell'


type ProjectCardPropType = { 
    queryData: DBQueryData
}

const Table = ({queryData}: ProjectCardPropType) => {

    const data = React.useMemo(
        () => queryData.rows,
        [queryData]
    )

    const columns = React.useMemo(
        () => queryData.columns.map((col) => ({
            Header: col,
            accessor: col, 
        })),
        [queryData]
    )

    const defaultColumn = {
        Cell: EditableCell,
    }      

    const {
        getTableProps,
        getTableBodyProps,
        headerGroups,
        rows,
        prepareRow,
      } = useTable({ columns, data, defaultColumn })
    

    return (
        <React.Fragment>
            <table {...getTableProps()} className={"table is-bordered is-striped is-narrow is-hoverable is-fullwidth"}>
                <thead>
                    {headerGroups.map(headerGroup => (
                        <tr {...headerGroup.getHeaderGroupProps()}>
                            {headerGroup.headers.map(column => (
                                <th {...column.getHeaderProps()}>
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
                            <tr {...row.getRowProps()}>
                                {row.cells.map(cell => {
                                    return (<td {...cell.getCellProps()}>
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