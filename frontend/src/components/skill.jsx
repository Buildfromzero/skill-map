export default function Skill({ id, name }) {
    return (
        <tr>
            <th scope="row">{id}</th>
            <td>{name}</td>
        </tr>
    );
}