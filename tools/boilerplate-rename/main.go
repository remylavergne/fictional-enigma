package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	oldPackage := "co.touchlab.kampkit"
	newPackage := "be.afelio.boilerplate"

	replacePackageOnFiles(getPackagePaths(), oldPackage, newPackage)
	moveFilesToNewPackage(getPathsToRename(), oldPackage, newPackage)

}

func up(step int) string {
	steps := ""
	for i := 0; i < step; i++ {
		steps = steps + "../"
	}
	return steps
}

func replacePackageOnFiles(paths []string, oldPackage string, newPackage string) {

	for _, path := range paths {
		read, err := ioutil.ReadFile(up(2) + path)
		if err != nil {
			panic(err)
		}

		// fmt.Println(path)

		newContents := strings.Replace(string(read), oldPackage, newPackage, -1)

		err = ioutil.WriteFile(up(2)+path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Done!")
}

func updatePath(currentPath string, oldPackage string, newPackage string) string {
	return strings.Replace(currentPath, oldPackage, newPackage, 1)
}

func moveFilesToNewPackage(paths []string, oldPackage string, newPackage string) {

	for _, path := range paths {
		// Receive => app/src/main/java/co/touchlab/kampkit
		updatedPath := updatePath(path, oldPackage, newPackage)
		err := os.MkdirAll(up(2)+updatedPath, 0755)

		if err != nil {
			panic(err)
		}

		errr := filepath.Walk(up(2)+path,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				// Ignore top level path
				if path == up(2)+path {
					return nil
				}

				// Change path
				newPath := strings.Replace(path, oldPackage, newPackage, 1)
				fmt.Println("New path =>", newPath)

				if info.IsDir() {
					err := os.Mkdir(newPath, 0755)
					if err != nil {
						panic(err)
					}
				} else {
					read, err := ioutil.ReadFile(path)
					if err != nil {
						panic(err)
					}

					err = ioutil.WriteFile(newPath, []byte(read), 0755)
					if err != nil {
						panic(err)
					}
				}

				return nil
			})
		if errr != nil {
			log.Println(err)
		}
	}
}

func getPackagePaths() []string {
	return []string{
		"app/build.gradle.kts",
		"app/src/main/AndroidManifest.xml",
		"app/src/main/java/co/touchlab/kampkit/android/BreedViewModel.kt",
		"app/src/main/java/co/touchlab/kampkit/android/MainActivity.kt",
		"app/src/main/java/co/touchlab/kampkit/android/MainApp.kt",
		"app/src/main/java/co/touchlab/kampkit/android/adapter/MainAdapter.kt",
		"app/src/main/java/co/touchlab/kampkit/android/adapter/MainViewHolder.kt",
		"app/src/main/res/layout/activity_main.xml",
		"ios/KaMPKitiOS.xcodeproj/project.pbxproj",
		"shared/build.gradle.kts",
		"shared/src/androidMain/kotlin/co/touchlab/kampkit/KoinAndroid.kt",
		"shared/src/androidMain/kotlin/co/touchlab/kampkit/PlatformAndroid.kt",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit/BaseTest.kt",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit/CoroutineTestRule.kt",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit/KoinTest.kt",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit/TestUtilAndroid.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/AppInfo.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/DatabaseHelper.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/Koin.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/Platform.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/ktor/DogApiImpl.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/ktor/KtorApi.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/models/BreedModel.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/models/DataState.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/response/BreedResult.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/sqldelight/CoroutinesExtensions.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/BaseTest.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/BreedModelTest.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/ConcurrencyTest.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/SqlDelightTest.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/TestAppInfo.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/TestUtil.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/mock/ClockMock.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/mock/KtorApiMock.kt",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit/KoinIOS.kt",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit/MainScope.kt",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit/NativeViewModel.kt",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit/PlatformiOS.kt",
		"shared/src/iosTest/kotlin/co/touchlab/kampkit/BaseTest.kt",
		"shared/src/iosTest/kotlin/co/touchlab/kampkit/KoinTest.kt",
		"shared/src/iosTest/kotlin/co/touchlab/kampkit/TestUtilIOS.kt",
		"shared/src/main/AndroidManifest.xml",
	}
}

func getPathsToRename() []string {
	return []string{
		"app/src/main/java/co/touchlab/kampkit",
		"shared/src/androidMain/kotlin/co/touchlab/kampkit",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit",
		"shared/src/commonMain/sqldelight/co/touchlab/kampkit",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit",
		"shared/src/iosTest/kotlin/co/touchlab/kampkit",
	}
}
