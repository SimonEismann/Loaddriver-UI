<script>
  import Panel from "../../components/Panel.svelte";
  import NumberInput from "../../components/form/NumberInput.svelte";
  import TextInput from "../../components/form/TextInput.svelte";
  import IntensityChart from "../../components/IntensityChart.svelte";
  import Button from "../../components/form/Button.svelte";
  import { uploadIntensityFile } from "../../model/api.js";
  import { Link } from "svelte-routing";
  import { Point2D } from "../../model/intensity-data.js";
  import { navigate } from "svelte-routing";
  let duration = 60;
  let amplitude = 10;
  let min = 0;
  let numOfPeriods = 1;
  let filename = "";
  let previewData = [];

  $: {
    previewData = [];
    const stretch = duration / (2.0 * Math.PI);
    for (let i = 0; i <= duration; i++) {
      previewData.push(
        new Point2D(
          i,
          amplitude - min + amplitude * -Math.cos((i / stretch) * numOfPeriods)
        )
      );
    }
  }

  const upload = async () => {
    try {
      await uploadIntensityFile(
        filename,
        previewData.map(point => `${point.x + 0.5},${point.y}`).join("\n")
      );
      navigate("/intensities");
    } catch (error) {
      console.log(error);
    }
  };
</script>

<Panel
  title="Sine Intensity Wizard"
  subtitle="Configure the intensity file and upload it directly to the server to
  make it available for experiments">
  <form on:submit|preventDefault={upload}>
    <TextInput
      label="Filename (Required)"
      bind:value={filename}
      required="true" />
    <NumberInput label="Duration" bind:value={duration} />
    <NumberInput label="Number of periods" bind:value={numOfPeriods} />
    <NumberInput label="Amplitude" bind:value={amplitude} />
    <NumberInput label="Minimum" bind:value={min} />
    <h2 style="margin-bottom: 1em; margin-top: 1em">Preview</h2>
    <IntensityChart data={previewData} />
    <Button
      backgroundColor="var(--primary-action-color)"
      type="submit"
      value="Generate"
      title="Generate and upload to server"
      icon="fa-cogs" />
    <Link to="/intensities">
      <Button
        backgroundColor="red"
        type="button"
        value="Cancel"
        title="Cancel and return to type selection" />
    </Link>
  </form>
</Panel>
