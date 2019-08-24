cmake_minimum_required(VERSION 3.0.0)
project(language VERSION 0.1.0)

include(CTest)
enable_testing()

#add_executable(language english.c german.c)
#add_library(language SHARED english.c german.c)
add_library(language STATIC english.c german.c)
set(CPACK_PROJECT_NAME ${PROJECT_NAME})
set(CPACK_PROJECT_VERSION ${PROJECT_VERSION})
include(CPack)