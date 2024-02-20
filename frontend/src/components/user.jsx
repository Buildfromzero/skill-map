import { Link } from 'react-router-dom';


export default function User({ id, fullName, email }) {
  return (
    <tr>
      <th scope="row">
        <Link to={`/users/${id}`}>{id}</Link>
      </th>
      <td>

        <Link to={`/users/${id}`}>{fullName}</Link>
      </td>
      <td>{email}</td>
      <td>
        <span class="badge rounded-pill text-bg-primary">Go</span>
        <span class="badge rounded-pill text-bg-secondary">Python</span>
        <span class="badge rounded-pill text-bg-success">Rust</span>
        <span class="badge rounded-pill text-bg-danger">C</span>
      </td>
    </tr>
  );
}