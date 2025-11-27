import { FolderPlus } from "lucide-react";

export default function AddProjectCard({
  onSelect,
}: {
  onSelect: (path: string) => void;
}) {
  function handlePick() {
    const input = document.createElement("input");
    input.type = "file";
    input.multiple = true;

    // enable folder picker (cross browser)
    (input as any).webkitdirectory = true;
    (input as any).mozdirectory = true;
    (input as any).msdirectory = true;
    (input as any).odirectory = true;

    input.onchange = () => {
      const files = Array.from(input.files || []);
      if (files.length > 0) {
        const first = files[0] as any;
        const relative = first.webkitRelativePath || first.name;
        const topFolder = relative.split("/")[0];
        onSelect(topFolder);
      }
    };

    input.click();
  }

  return (
    <button
      onClick={handlePick}
      className="
        flex flex-col items-center justify-center
        border-2 border-dashed border-indigo-300
        rounded-2xl p-6 w-full h-40
        text-indigo-500 hover:border-indigo-500 hover:text-indigo-700
        transition cursor-pointer
      "
    >
      <FolderPlus size={36} />
      <span className="mt-3 font-medium text-sm">Add Project</span>
    </button>
  );
}
