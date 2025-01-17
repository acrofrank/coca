package app

import (
	"encoding/json"
	"github.com/modernizing/coca/cmd/cmd_util"
	"github.com/modernizing/coca/pkg/adapter/cocafile"
	"github.com/modernizing/coca/pkg/application/analysis"
	"github.com/modernizing/coca/pkg/application/analysis/tsapp"
	"github.com/modernizing/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
)

type AnalysisCmdConfig struct {
	Path        string
	ForceUpdate bool
	Lang        string
}

var (
	analysisCmdConfig AnalysisCmdConfig
)

var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "analysis code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var outputName string
		var ds []core_domain.CodeDataStruct

		ds = analysis.CommonAnalysis(output, analysisCmdConfig.Path, new(tsapp.TypeScriptIdentApp), cocafile.GoFileFilter, true)
		outputName = "pydeps.json"

		cModel, _ := json.MarshalIndent(ds, "", "\t")
		cmd_util.WriteToCocaFile(outputName, string(cModel))
	},
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Path, "path", "p", ".", "example -p core/main")
	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Lang, "lang", "l", "java", "example coca analysis -l java, typescript, python")
	analysisCmd.PersistentFlags().BoolVarP(&analysisCmdConfig.ForceUpdate, "force", "f", false, "force update -f")
}
