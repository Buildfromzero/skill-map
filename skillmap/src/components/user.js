export default function User({id,firstName,lastName}) {
  return (
    <tr>
        <th scope="row">{id}</th>
        <td>{firstName}</td>
        <td>{lastName}</td>
        <td>
          <span class="badge rounded-pill text-bg-primary">Go</span>
          <span class="badge rounded-pill text-bg-secondary">Python</span>
          <span class="badge rounded-pill text-bg-success">Rust</span>
          <span class="badge rounded-pill text-bg-danger">C</span>
        </td>
    </tr>
  );
}