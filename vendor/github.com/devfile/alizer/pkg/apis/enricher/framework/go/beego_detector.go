/*******************************************************************************
 * Copyright (c) 2022 Red Hat, Inc.
 * Distributed under license by Red Hat, Inc. All rights reserved.
 * This program is made available under the terms of the
 * Eclipse Public License v2.0 which accompanies this distribution,
 * and is available at http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 * Red Hat, Inc.
 ******************************************************************************/

package enricher

import (
	"context"
	"regexp"

	"github.com/devfile/alizer/pkg/apis/model"
	"github.com/devfile/alizer/pkg/utils"
	"golang.org/x/mod/modfile"
)

type BeegoDetector struct{}

func (b BeegoDetector) GetSupportedFrameworks() []string {
	return []string{"Beego"}
}

func (b BeegoDetector) GetApplicationFileInfos(componentPath string, ctx *context.Context) []model.ApplicationFileInfo {
	return []model.ApplicationFileInfo{
		{
			Context: ctx,
			Root:    componentPath,
			Dir:     "conf",
			File:    "app.conf",
		},
	}
}

// DoFrameworkDetection uses a tag to check for the framework name
func (b BeegoDetector) DoFrameworkDetection(language *model.Language, goMod *modfile.File) {
	if hasFramework(goMod.Require, "github.com/beego/beego") {
		language.Frameworks = append(language.Frameworks, "Beego")
	}
}

type ApplicationPropertiesFile struct {
	Dir  string
	File string
}

// DoPortsDetection searches for the port in conf/app.conf
func (b BeegoDetector) DoPortsDetection(component *model.Component, ctx *context.Context) {
	fileContents, err := utils.GetApplicationFileContents(b.GetApplicationFileInfos(component.Path, ctx))
	if err != nil {
		return
	}
	for _, fileContent := range fileContents {
		re := regexp.MustCompile(`httpport\s*=\s*(\d+)`)
		ports := utils.FindAllPortsSubmatch(re, fileContent, 1)
		if len(ports) > 0 {
			component.Ports = ports
			return
		}
	}
}
