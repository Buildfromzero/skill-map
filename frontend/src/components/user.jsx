export default function User({ id, fullName, email }) {
  return (
    <tr>
      <th scope="row">{id}</th>
      <td>{fullName}</td>
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