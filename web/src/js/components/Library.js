import React, { Component } from 'react'
import PropTypes from 'prop-types'

class Library extends Component {
  render() {
    return (
      <div className="library">
        <List rows={[ ["title", "info"],
                      ["cowboy bebop", "fav"],
                      ["penguindrum", "??? what even is"],
                      ["psycho-pass", "I still need to finish this xD"] ]} />
      </div>
    );
  }
}


const List = ({ rows }) => (
  <div className="list">
    {rows.map((row, i) => (
      <ListRow key={i} columns={row} />
    ))}
  </div>
)

const ListRow = ({ columns }) => (
  <div className="row">
    {columns.map((column, i) => (
      <ListRowColumn key={i} text={column} />
    ))}
  </div>
)

const ListRowColumn = ({ text }) => (
    <div className="column">
      {text}
    </div>
)

List.propTypes = {
  rows: PropTypes.array
}
ListRow.propTypes = {
  columns: PropTypes.array
}
ListRowColumn.propTypes = {
  text: PropTypes.string
}

export default Library
