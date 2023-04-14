import styles from './cheatsheet.module.scss'
import React, { useRef, useState } from 'react'
import { DBConnType } from '../../../data/defaults'
import CheatsheetCommand from './command'
import _ from 'lodash'
import lunr from 'lunr'

type CheatSheetPropType = {
    dbType: DBConnType
    onClose: () => void
}

const CheatSheetModal = ({ dbType, onClose }: CheatSheetPropType) => {

    const [search, setSearch] = useState<lunr.Index.Result[] | undefined>()

    const searchInputRef = useRef<HTMLInputElement | null>(null)

    const cmdList = dbType === DBConnType.POSTGRES ? pgcommandList : dbType === DBConnType.MYSQL ? mysqlcommandList : mongoCommandList

    const searchIndex = lunr(function () {
        this.field('title', { boost: 10 })
        this.field('description')
        this.metadataWhitelist.push("where")
        cmdList.forEach((cmd, i) => {
            this.add({ ...cmd, id: i })
        })
    })

    const doSearch = () => {
        const queryString = searchInputRef.current!.value
        if (queryString === "") {
            setSearch(undefined)
            return
        }
        setSearch(searchIndex.search(queryString))
    }

    return (<React.Fragment>
        <div className="modal is-active">
            <div className="modal-background"></div>
            <div className="modal-card">
                <header className="modal-card-head">
                    <p className="modal-card-title"><i className="fas fa-book"></i>&nbsp;{_.capitalize(dbType)} Cheatsheet</p>
                    <button className="delete" aria-label="close" onClick={onClose}></button>
                </header>
                <section className="modal-card-body">
                    <div className="field">
                        <div className="control has-icons-left">
                            <input
                                ref={searchInputRef}
                                className="input"
                                type="text"
                                placeholder="Find a command"
                                onChange={doSearch} />
                            <span className="icon is-small is-left">
                                <i className="fas fa-magnifying-glass"></i>
                            </span>
                        </div>
                    </div>
                    <br />
                    {!search && (cmdList).map((cmd, i) => {
                        return <CheatsheetCommand cmd={cmd} isLast={i === cmdList.length - 1} />
                    })}
                    {search && search.map((result, i) => {
                        return <CheatsheetCommand cmd={cmdList[parseInt(result.ref)]} isLast={i === search.length - 1} />
                    })}
                    {search && search.length === 0 && <>
                        <h2 style={{ textAlign: 'center' }}>No commands found</h2>
                    </>}
                </section>
                <footer className="modal-card-foot">
                    <button className="button" onClick={onClose}>Close</button>
                </footer>
            </div>
        </div>

    </React.Fragment>)
}


export default CheatSheetModal


const mysqlcommandList = [
    {
        "title": "Select all",
        "description": "Gets all the columns and rows from the table.",
        "command": "SELECT * FROM <table name>;"
    },
    {
        "title": "Select only given column names",
        "description": "Gets only the given columns and all rows from the table.",
        "command": "SELECT <column name 1>, <column name 2>, <column name 3> FROM <table name>;"
    },
    {
        "title": "Select with column alias",
        "description": "Gets only the given columns and all rows from the table.",
        "command": "SELECT <column name> AS <alias name> FROM <table name>;"
    },
    {
        "title": "Select Distinct",
        "description": "Gets all rows with unique values in given column from the table",
        "command": "SELECT DISTINCT <column name> FROM <table name>;"
    },
    {
        "title": "Select with limit",
        "description": "Gets all the columns from the table with first n number of rows",
        "command": "SELECT * FROM <table name> LIMIT <number>;"
    },
    {
        "title": "Select with limit and offset",
        "description": "Gets all the columns and n number of rows starting from offset from the table.",
        "command": "SELECT * FROM <table name> LIMIT <limit number> OFFSET <offset number>;"
    },
    {
        "title": "Select with ordering",
        "description": "Gets all columns and rows from the table in sorted order according to given sort expression. Sort name can be a column or an expression.",
        "command": "SELECT * FROM <table name> ORDER BY <sort name> [ASC | DESC];"
    },
    {
        "title": "Select using where clause with = operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches the given value.",
        "command": "SELECT * FROM <table name> WHERE <column name> = <value>;"
    },
    {
        "title": "Select using where clause with != operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value does not matche the given value.",
        "command": "SELECT * FROM <table name> WHERE <column name> != <value>;"
    },
    {
        "title": "Select using where clause with AND operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches the given value and the other column names match other given values.",
        "command": "SELECT * FROM <table name> WHERE <column name 1> = <value 1> AND <column name 2> = <value 2>;"
    },
    {
        "title": "Select using where clause with OR operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches the given value or the other column names match other given values.",
        "command": "SELECT * FROM <table name> WHERE <column name 1> = <value 1> OR <column name 2> = <value 2>;"
    },
    {
        "title": "Select using where clause with IN operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches any value in the given list",
        "command": "SELECT * FROM <table name> WHERE <column name> IN (<value 1>, <value 2>, <value 3>);"
    },
    {
        "title": "Select using where clause with LIKE operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches the specified pattern",
        "command": "SELECT * FROM <table name> WHERE <column name> LIKE '<pattern with % for wildcard and _ for single char>';"
    },
    {
        "title": "Select using where clause with NOT LIKE operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value does not match the specified pattern",
        "command": "SELECT * FROM <table name> WHERE <column name> NOT LIKE '<pattern with % as wildcard and _ for single char>';"
    },
    {
        "title": "Select using where clause with BETWEEN operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value is in the range of specified values",
        "command": "SELECT * FROM <table name> WHERE <column name> BETWEEN <value 1> AND <value 2>;"
    },
    {
        "title": "Select using where clause with IS NULL operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value is equal to null.",
        "command": "SELECT * FROM <table name> WHERE <column name> IS NULL;"
    },
    {
        "title": "Select using where clause with IS NOT NULL operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value is not equal to null.",
        "command": "SELECT * FROM <table name> WHERE <column name> IS NOT NULL;"
    },
    {
        "title": "Using Inner join",
        "description": "Inner joins a table with other table on given condition. If the given condition is satified, the inner join creates a new row that contains columns from both tables and adds this new row to the result.",
        "command": "SELECT <columns> FROM <table name 1> INNER JOIN <table name 2> ON <table 1 column name> = <table 2 column name>;"
    },
    {
        "title": "Using Left join",
        "description": "Left joins a table with other table on given condition. If the given condition is satified, the left join creates a new row that contains columns from both tables and adds this new row to the result. If condition is not matched, it adds a new row with values from left table.",
        "command": "SELECT <columns> FROM <table name 1> LEFT JOIN <table name 2> ON <table 1 column name> = <table 2 column name>;"
    },
    {
        "title": "Using Right join",
        "description": "Right joins a table with other table on given condition. If the given condition is satified, the right join creates a new row that contains columns from both tables and adds this new row to the result. If condition is not matched, it adds a new row with values from right table.",
        "command": "SELECT <columns> FROM <table name 1> RIGHT JOIN <table name 2> ON <table 1 column name> = <table 2 column name>;"
    },
    {
        "title": "Using Cross join",
        "description": "A cross join clause allows you to produce a Cartesian Product of rows in two or more tables.",
        "command": "SELECT <columns> FROM <table name 1> CROSS JOIN <table name 2>;"
    },
    {
        "title": "Using Group by and aggregate",
        "description": "Divides the rows returned by select into groups and for each group applies the given aggregate function.",
        "command": "SELECT <column name>, <aggregate function> FROM <table name> GROUP BY <column name>;"
    },
    {
        "title": "Using having clause",
        "description": "The HAVING clause specifies a search condition for a group or an aggregate. ",
        "command": "SELECT <column name>, <aggregate function> FROM <table name> GROUP BY <column name> HAVING <condition>;"
    },
    {
        "title": "Inserting row into table",
        "description": "Inserts a rows with the given values for the given columns into the table.",
        "command": "INSERT INTO <table name> (<column 1>, <column 2>,...) VALUES (<value 1>, <value 2>,...);"
    },
    {
        "title": "Inserting multiple row into table",
        "description": "Inserts multiple rows with the given values for the given columns into the table.",
        "command": "INSERT INTO <table name> (<column 1>, <column 2>,...) VALUES (<value x1>, <value x2>,...), (<value y1>, <value y2>,...);"
    },
    {
        "title": "Create new table",
        "description": "Creates new table with the given table name and given columns and thier data types.",
        "command": "CREATE TABLE [IF NOT EXISTS] <table name> (<column name 1> <data type> <contraint>, <column name 2> <data type> <contraint>,...);"
    },
    {
        "title": "Rename a table",
        "description": "Renames an existing table to the new given table name.",
        "command": "ALTER TABLE [IF EXISTS] <table name> RENAME TO <new table name>;"
    },
    {
        "title": "Add a new column to existing table",
        "description": "Creates a new column with the given name and data type in a given existing table.",
        "command": "ALTER TABLE <table name> ADD COLUMN <column name> <data type> <constraint>;"
    },
    {
        "title": "Delete all data from given table",
        "description": "Removes all the data from the given table.",
        "command": "TRUNCATE TABLE <table name> [CASCADE];"
    },
]

const pgcommandList = [
    {
        "title": "Select all",
        "description": "Gets all the columns and rows from the table.",
        "command": "SELECT * FROM <table name>;"
    },
    {
        "title": "Select only given column names",
        "description": "Gets only the given columns and all rows from the table.",
        "command": "SELECT <column name 1>, <column name 2>, <column name 3> FROM <table name>;"
    },
    {
        "title": "Select with column alias",
        "description": "Gets only the given columns and all rows from the table.",
        "command": "SELECT <column name> AS <alias name> FROM <table name>;"
    },
    {
        "title": "Select Distinct",
        "description": "Gets all rows with unique values in given column from the table",
        "command": "SELECT DISTINCT <column name> FROM <table name>;"
    },
    {
        "title": "Select with limit",
        "description": "Gets all the columns from the table with first n number of rows",
        "command": "SELECT * FROM <table name> LIMIT <number>;"
    },
    {
        "title": "Select with limit and offset",
        "description": "Gets all the columns and n number of rows starting from offset from the table.",
        "command": "SELECT * FROM <table name> LIMIT <limit number> OFFSET <offset number>;"
    },
    {
        "title": "Select with ordering",
        "description": "Gets all columns and rows from the table in sorted order according to given sort expression. Sort name can be a column or an expression.",
        "command": "SELECT * FROM <table name> ORDER BY <sort name> [ASC | DESC];"
    },
    {
        "title": "Select using where clause with = operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches the given value.",
        "command": "SELECT * FROM <table name> WHERE <column name> = <value>;"
    },
    {
        "title": "Select using where clause with != operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value does not matche the given value.",
        "command": "SELECT * FROM <table name> WHERE <column name> != <value>;"
    },
    {
        "title": "Select using where clause with AND operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches the given value and the other column names match other given values.",
        "command": "SELECT * FROM <table name> WHERE <column name 1> = <value 1> AND <column name 2> = <value 2>;"
    },
    {
        "title": "Select using where clause with OR operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches the given value or the other column names match other given values.",
        "command": "SELECT * FROM <table name> WHERE <column name 1> = <value 1> OR <column name 2> = <value 2>;"
    },
    {
        "title": "Select using where clause with IN operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches any value in the given list",
        "command": "SELECT * FROM <table name> WHERE <column name> IN (<value 1>, <value 2>, <value 3>);"
    },
    {
        "title": "Select using where clause with LIKE operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches the specified pattern",
        "command": "SELECT * FROM <table name> WHERE <column name> LIKE '<pattern with % for wildcard and _ for single char>';"
    },
    {
        "title": "Select using where clause with NOT LIKE operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value does not match the specified pattern",
        "command": "SELECT * FROM <table name> WHERE <column name> NOT LIKE '<pattern with % as wildcard and _ for single char>';"
    },
    {
        "title": "Select using where clause with ILIKE operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value matches the specified pattern case-insensitively",
        "command": "SELECT * FROM <table name> WHERE <column name> ILIKE '<pattern with % as wildcard and _ for single char>';"
    },
    {
        "title": "Select using where clause with BETWEEN operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value is in the range of specified values",
        "command": "SELECT * FROM <table name> WHERE <column name> BETWEEN <value 1> AND <value 2>;"
    },
    {
        "title": "Select using where clause with IS NULL operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value is equal to null.",
        "command": "SELECT * FROM <table name> WHERE <column name> IS NULL;"
    },
    {
        "title": "Select using where clause with IS NOT NULL operator",
        "description": "Gets all the columns from the table and all the rows where the given column name value is not equal to null.",
        "command": "SELECT * FROM <table name> WHERE <column name> IS NOT NULL;"
    },
    {
        "title": "Using Inner join",
        "description": "Inner joins a table with other table on given condition. If the given condition is satified, the inner join creates a new row that contains columns from both tables and adds this new row to the result.",
        "command": "SELECT <columns> FROM <table name 1> INNER JOIN <table name 2> ON <table 1 column name> = <table 2 column name>;"
    },
    {
        "title": "Using Left join",
        "description": "Left joins a table with other table on given condition. If the given condition is satified, the left join creates a new row that contains columns from both tables and adds this new row to the result. If condition is not matched, it adds a new row with values from left table.",
        "command": "SELECT <columns> FROM <table name 1> LEFT JOIN <table name 2> ON <table 1 column name> = <table 2 column name>;"
    },
    {
        "title": "Using Right join",
        "description": "Right joins a table with other table on given condition. If the given condition is satified, the right join creates a new row that contains columns from both tables and adds this new row to the result. If condition is not matched, it adds a new row with values from right table.",
        "command": "SELECT <columns> FROM <table name 1> RIGHT JOIN <table name 2> ON <table 1 column name> = <table 2 column name>;"
    },
    {
        "title": "Using Full outer join",
        "description": "Full outer joins a table with other table on given condition. The full outer join or full join returns a result that contains all rows from both left and right tables, with the matching rows from both sides if available. In case there is no match, the columns of the table will be filled with NULL.",
        "command": "SELECT <columns> FROM <table name 1> FULL OUTHER JOIN <table name 2> ON <table 1 column name> = <table 2 column name>;"
    },
    {
        "title": "Using Cross join",
        "description": "A cross join clause allows you to produce a Cartesian Product of rows in two or more tables.",
        "command": "SELECT <columns> FROM <table name 1> CROSS JOIN <table name 2>;"
    },
    {
        "title": "Using Natural join",
        "description": "A natural join is a join that creates an implicit join based on the same column names in the joined tables.",
        "command": "SELECT <columns> FROM <table name 1> NATURAL [INNER, LEFT, RIGHT] JOIN <table name 2>;"
    },
    {
        "title": "Using Group by and aggregate",
        "description": "Divides the rows returned by select into groups and for each group applies the given aggregate function.",
        "command": "SELECT <column name>, <aggregate function> FROM <table name> GROUP BY <column name>;"
    },
    {
        "title": "Using having clause",
        "description": "The HAVING clause specifies a search condition for a group or an aggregate. ",
        "command": "SELECT <column name>, <aggregate function> FROM <table name> GROUP BY <column name> HAVING <condition>;"
    },
    {
        "title": "Inserting row into table",
        "description": "Inserts a rows with the given values for the given columns into the table.",
        "command": "INSERT INTO <table name> (<column 1>, <column 2>,...) VALUES (<value 1>, <value 2>,...);"
    },
    {
        "title": "Inserting multiple row into table",
        "description": "Inserts multiple rows with the given values for the given columns into the table.",
        "command": "INSERT INTO <table name> (<column 1>, <column 2>,...) VALUES (<value x1>, <value x2>,...), (<value y1>, <value y2>,...);"
    },
    {
        "title": "Updating values in the table",
        "description": "Updates the given value for the column name in the rows where the given condition matches.",
        "command": "UPDATE <table name> SET <column name> = <value> WHERE <condition>;"
    },
    {
        "title": "Deleting rows from table",
        "description": "Deletes rows from given table where the given condition matches.",
        "command": "DELETE FROM <table name> WHERE <condition>;"
    },
    {
        "title": "Create new table",
        "description": "Creates new table with the given table name and given columns and thier data types.",
        "command": "CREATE TABLE [IF NOT EXISTS] <table name> (<column name 1> <data type> <contraint>, <column name 2> <data type> <contraint>,...);"
    },
    {
        "title": "Rename a table",
        "description": "Renames an existing table to the new given table name.",
        "command": "ALTER TABLE [IF EXISTS] <table name> RENAME TO <new table name>;"
    },
    {
        "title": "Add a new column to existing table",
        "description": "Creates a new column with the given name and data type in a given existing table.",
        "command": "ALTER TABLE <table name> ADD COLUMN <column name> <data type> <constraint>;"
    },
    {
        "title": "Delete all data from given table",
        "description": "Removes all the data from the given table.",
        "command": "TRUNCATE TABLE <table name> [CASCADE];"
    },
]

const mongoCommandList = [
    {
        "title": "Find all",
        "description": "Gets all the documents from the collection.",
        "command": "db.<collection name>.find();"
    },
    {
        "title": "Find with condition",
        "description": "Gets all the documents matching the filter from the collection.",
        "command": "db.<collection name>.find({<filter>});"
    },
    {
        "title": "Find with limit",
        "description": "Gets first n documents from the collection.",
        "command": "db.<collection name>.find().limit(<number>);"
    },
    {
        "title": "Find one document",
        "description": "Gets the first document from the collection matching the filter.",
        "command": "db.<collection name>.findOne({<filter>});"
    },
    {
        "title": "Insert one document",
        "description": "Inserts one document into the collection.",
        "command": "db.<collection name>.insertOne({<document>});"
    },
    {
        "title": "Insert documents",
        "description": "Inserts one or many documents into the collection.",
        "command": "db.<collection name>.insert([{<document 1>},{<document 2>}, ... ]);"
    },
    {
        "title": "Delete one document",
        "description": "Deletes one document from the collection matching the filter.",
        "command": "db.<collection name>.deleteOne({<filter>});"
    },
    {
        "title": "Delete documents",
        "description": "Deletes one or many documents from the collection matching the filter.",
        "command": "db.<collection name>.deleteMany({<filter>});"
    },
    {
        "title": "Update one document",
        "description": "Updates one document from the collection matching the filter.",
        "command": "db.<collection name>.updateOne({<filter>}, {<update>});"
    },
    {
        "title": "Update documents",
        "description": "Updates one or many documents from the collection matching the filter.",
        "command": "db.<collection name>.updateMany({<filter>}, {<update>});"
    },
    {
        "title": "Replace one document",
        "description": "Completely replaces one document matching the filter in the collection with given document.",
        "command": "db.<collection name>.replaceOne({<filter>}, {<document>});"
    },
    {
        "title": "Aggregate",
        "description": "Processess the given aggregation pipeline",
        "command": "db.<collection name>.aggregate([{<stage 1>}, {<stage 2>}, ...]);"
    },
]