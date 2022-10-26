<script>
  export let incidents = []
</script>

<h1>Incident Index</h1>

<table border="1" cellpadding="10">
  {#if incidents.length > 0}
    <thead>
      {#each Object.keys(incidents[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each incidents as incident}
    <tr>
      {#each Object.keys(incident) as key}
        {#if key.toLowerCase() === "id"}
          <td><a href={`/${incident.id || 0}`}>{incident[key]}</a></td>
        {:else}
          <td>{incident[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<br />

<a href={`/new`}>New Incident</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
