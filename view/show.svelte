<script>
  export let incident = {}
</script>

<h1>Show Incident</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(incident) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(incident) as key}
      <td>{incident[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href={`/`}>Back</a>
<span> | </span>
<a href={`/${incident.id || 0}/edit`}>Edit</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
