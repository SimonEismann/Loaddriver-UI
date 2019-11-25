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
  let min = 0;
  let peak = 10;
  let numOfPeaks = 1;
  let filename = "";
  let previewData = [];

  $: {
    previewData = [];
    const stepSize = duration / (numOfPeaks * 2.0);
    let m = (peak - min) / stepSize;
    let previous = min;
    previewData.push(new Point2D(0, min));
    for (let i = 1; i <= duration; i++) {
      let next = previous + m;
      previewData.push(new Point2D(i, next < 0 ? 0 : next));
      previous = next < 0 ? 0 : next;
      if (i % stepSize == 0) {
        m *= -1.0;
      }
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
  title="Peak Intensity Wizard"
  subtitle="Configure the intensity file and upload it directly to the server to
  make it available for experiments">
  <form on:submit|preventDefault={upload}>
    <TextInput
      label="Filename (Required)"
      bind:value={filename}
      required="true" />
    <NumberInput label="Duration" bind:value={duration} />
    <NumberInput label="Minimum Intensity" bind:value={min} />
    <NumberInput label="Peak Intensity" bind:value={peak} />
    <NumberInput label="Number of Peaks" bind:value={numOfPeaks} />
    <h2 style="margin-bottom: 1em; margin-top: 1em">Preview</h2>
    <IntensityChart data={previewData} curved={false} />
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
