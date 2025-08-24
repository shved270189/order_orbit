import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [users, setUsers] = useState([])

  useEffect(() => {
    const fetchData = () => {
      fetch(`http://localhost:${import.meta.env.VITE_PORT}/users`)
        .then(response => response.text())
        .then(data => setUsers(JSON.parse(data)))
        .catch(error => console.error('Error fetching data:', error))
    }
    fetchData()
  }, [])

  return (
    <div className="row">
      {users.length === 0 && (
        <div>
          <h2>Loading...</h2>
        </div>
      )}
      {users.length > 0 && (
        <table className="table table-striped">
          <thead>
            <tr>
              <th scope="col">ID</th>
              <th scope="col">Full Name</th>
              <th scope="col">Login</th>
              <th scope="col">Updated At</th>
              <th scope="col">Created At</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => {
              return (
                <tr key={user.id}>
                  <th scope="row">{user.id}</th>
                  <td>{user.full_name}</td>
                  <td>{user.login}</td>
                  <td>{user.updated_at}</td>
                  <td>{user.created_at}</td>
                </tr>
              )
            })}
          </tbody>
        </table>
      )}
    </div>
  )
}

export default App
