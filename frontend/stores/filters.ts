import { defineStore } from "pinia";

interface Selection {
  id?: string;
  label?: string;
}

interface FiltersState {
  selectedTags: Selection[];
  selectedCollection: Selection;
  textInput: string;
}

export const useFiltersStore = defineStore("filters", {
  state: (): FiltersState => {
    return {
      selectedTags: [],
      selectedCollection: {},
      textInput: "",
    };
  },
});
