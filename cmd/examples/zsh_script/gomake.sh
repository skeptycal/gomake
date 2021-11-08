#!/usr/bin/env zsh
# -*- coding: utf-8 -*-
    # shellcheck shell=bash
    # shellcheck source=/dev/null
    # shellcheck disable=2178,2128,2206,2034
#? -----------------------------> .repo_tools - tools for repo management for macOS with zsh
	#*	system functions
	#*  tested on macOS Big Sur and zsh 5.8
	#*	copyright (c) 2021 Michael Treanor
	#*	MIT License - https://www.github.com/skeptycal
#? -----------------------------> https://www.github.com/skeptycal
#? -----------------------------> parameter expansion tips
 #? ${PATH//:/\\n}    - replace all colons with newlines
 #? ${PATH// /}       - strip all spaces
 #? ${VAR##*/}        - return only final element in path (program name)
 #? ${VAR%/*}         - return only path (without program name)

#? -----------------------------> environment
    DOTFILES_INC=~/.dotfiles/zshrc_inc # TODO - this is getting lost somewhere in the mix

	# repo_tools.sh includes ansi_colors.sh
	. ${DOTFILES_INC}/repo_tools.sh || . $(which repo_tools.sh)

	declare -ix SET_DEBUG=${SET_DEBUG:-0}  		#! set to 1 for verbose testing
    dbecho "\${DOTFILES_INC}: ${DOTFILES_INC}"

    SCRIPT_NAME=${0##*/}
    dbecho "\${SCRIPT_NAME}: ${SCRIPT_NAME}"

#? -----------------------------> debug info
	_debug_tests() {
		if (( SET_DEBUG )); then
			printf '%b\n' "${WARN:-}Debug Mode Details for ${CANARY}${SCRIPT_NAME##*/}${RESET:-}"
			green "Today is $today"
			_test_args "$@"
			dbecho "DEBUG mode is enabled. Set it to 0 to disable these messages."
			dbecho ""
			dbecho "ANSI colors active: $(which ansi_colors.sh)"
			dbecho "Git repo tools active: $(which repo_tools.sh)"
		fi
		}
    _debug_tests "$@"

    dbinfo() {
        dbecho "$0: $@"
    }
#? -----------------------------> utilities
    SET_DEBUG=0
    _DEV_GOMAKE=0

    is_prod() { [ $SET_DEBUG -eq 0 ]; }
    is_dev() { [ $SET_DEBUG -ne 0 ]; }
    _go_version() {
        go version
        }

    version() {
        # echo $(git describe --tags $(git rev-list --tags --max-count=1))
        git tag | sort -V | tail -n 1
        }

    bump() {
        #* Bump version

        local old=$SET_DEBUG
        SET_DEBUG=0 # constant dev mode verbosity level

        local vv=
        local dev=

        vv=$(echo $(version) | cut -d '-' -f 1)
        [ -z $vv ] && vv='v0.1.0'
        dbecho "\$vv: $vv"

        local vv=${vv#v*}
        dbecho "\$vv (remove 'v'): $vv"

        local major=${vv%%.*}
        dbecho "\$major: $major"

        vv=${vv#*.}
        local minor=${vv%%.*}
        dbecho "\$minor: $minor"

        vv=${vv#*.}
        local patch=${vv%%.*}
        dbecho "\$patch: $patch"

        case "$1" in
            major)
                major=$(( major + 1 ))
                dbecho "\$major increased: $major"
                minor=0
                dbecho "\$minor reset: $minor"
                patch=0
                dbecho "\$patch reset: $patch"
                ;;

            minor)
                minor=$(( minor + 1 ))
                dbecho "\$minor increased: $minor"
                patch=0
                dbecho "\$patch reset: $patch"
                ;;

            patch)
                patch=$(( patch + 1 ))
                dbecho "\$patch increased: $patch"
                ;;

            dev)
                shift
                devtag "$@"
                return
                ;;

            *)
                echo "current version: $(version)"
                usage bump '[major|minor|patch|dev][message]'
                dbecho "default case \$version: $version"
                return 0
                ;;

        esac

        printf -v version "v%s.%s.%s" $major $minor $patch
        dbecho "\$version: $version"

        echo "new version: $version"
        git tag "$version"
        git push origin --tags
        git push origin --all
        SET_DEBUG=old # constant dev mode verbosity level

        }

    usage() {
        if [ -z "$1" ]; then
            white "usage: ${MAIN:-}${0} ${DARKGREEN:-}app [args]"
            return 1
        fi

        local app="$1"
        shift
        white "usage: ${MAIN:-}${app} ${DARKGREEN:-}${@}"
        }

    devtag() {
        #* Tag Dev version

        local vv=$(version)
        local dev=
        local message=

        if [ -n "$1" ]; then
            message="(GoBot) devtag version ${version}: ${1}"
            shift
            echo "$message"
        fi

        vv=${vv%%-*}
        [ -z $vv ] && vv='v0.1.0'
        dbecho "\$vv: $vv"

        printf -v dev "%16.16s" $(date +'%s%N')
        dbecho "\$dev: $dev"

        version="${vv}-${dev}"
        dbecho "\$version: $version"

        git tag "$version"
        echo $version >|${REPO_PATH}/${REPO_NAME}/VERSION
        git add VERSION
        if [ -n "$message" ]; then
            git commit -m "$message"
        else
            git commit -m "(GoBot) devtag version ${version}" # only used if no message was provided
        fi
        git push origin --tags
        git push origin --all

        white "new Git version tag: ${MAIN:-}$version"
        }

    clean_template_name() {
        if [ -z "$1" ]; then
            usage $0 '[files]'
            return 0
        fi
        for f in $@; do
            sed -i '' -e "s/gorepotemplate/${REPO_NAME}/g" ${f}
        done
        }

    template() {
        if [ -z "$1" ]; then
            files=( CODE_OF_CONDUCT.md LICENSE README.md SECURITY.md contributing.md go.doc go.test.sh idea.md .gitignore .editorconfig )
        else
            files=$@
        fi

        for f in $files; do
            cp -arf ${LOCAL_TEMPLATE_PATH}/${f} .
        done;

        clean_template_name $files
        }

#? -----------------------------> gomake setup
    _setup_variables() {
        #* set repo variables
		#* general information
			YEAR=$( date +'%Y'; )
            dbecho "\${YEAR}: ${YEAR}"

		#* local repo information
			REPO_PATH="${PWD%/*}"
			REPO_NAME="${PWD##*/}"
            LOCAL_USER_PATH="${GOPATH}/src/github.com/$(whoami)"
            LOCAL_TEMPLATE_PATH="${LOCAL_USER_PATH}/gorepotemplate"
            EXAMPLE_PATH="cmd/example/${REPO_NAME}"
            EXAMPLE_FILE="${EXAMPLE_PATH}/main.go"

            dbinfo "Setup Local Repo (${REPO_NAME})"
            dbinfo "\${REPO_PATH}: ${REPO_PATH}"
            dbecho "\${REPO_NAME}: ${REPO_NAME}"
            dbecho "\${LOCAL_USER_PATH}: ${LOCAL_USER_PATH}"
            dbecho "\${LOCAL_TEMPLATE_PATH}: ${LOCAL_TEMPLATE_PATH}"
            dbecho "\${EXAMPLE_PATH}: ${EXAMPLE_PATH}"
            dbecho "\${EXAMPLE_FILE}: ${EXAMPLE_FILE}"

		#* github repo information
            GITHUB_TEMPLATE_PATH="https://github.com/skeptycal/gorepotemplate"
            GITHUB_USERNAME="${_gh_user}"
            GITHUB_URL="https://github.com/${GITHUB_USERNAME}"
            GITHUB_REPO_URL="${GITHUB_URL}/${REPO_NAME}"
            GO_GET_URL="github.com/${GITHUB_USERNAME}/${REPO_NAME}"
            GITHUB_DOCS_URL="${GITHUB_REPO_URL}/docs"
            PAGES_URL="https://${GITHUB_USERNAME}.github.io/${REPO_NAME}"

            dbinfo "Setup Remote Repo (${REPO_NAME})"
            dbecho "\${GITHUB_TEMPLATE_PATH}: ${GITHUB_TEMPLATE_PATH}"
            dbecho "\${GITHUB_USERNAME}: ${GITHUB_USERNAME}"
            dbecho "\${GITHUB_URL}: ${GITHUB_URL}"
            dbecho "\${GITHUB_REPO_URL}: ${GITHUB_REPO_URL}"
            dbecho "\${GO_GET_URL}: ${GO_GET_URL}"
            dbecho "\${GITHUB_DOCS_URL}: ${GITHUB_DOCS_URL}"
            dbecho "\${PAGES_URL}: ${PAGES_URL}"

        #* file header blurbs
			BLURB_GO=$( _file_blurb )
			BLURB_INI=$( _file_blurb '#' )

            dbecho "\${BLURB_GO}: ${BLURB_GO}"
            dbecho "\${BLURB_INI}: ${BLURB_INI}"
    }
    _setup_environment() {
        dbinfo "Setup Environment"

        #/ ******* test setup
            #/ if the current directory is gomake_test, a special version
            #/ of this script will run for testing purposes
            #/ if PWD == gomake_test, clear test directory and remake everything ...
            if [[ ${PWD##*/} = "gomake_test" ]]; then
                warn "gomake_test directory found ... entering test mode."
                _DEV_GOMAKE=1
                cd ~
                rm -rf ~/gomake_test
                mkdir ~/gomake_test
                cd ~/gomake_test
            fi
        #/ ******* end test setup

        #* gh must be authenticated to use this script.
		_gh_auth_username
		if [[ -z ${_gh_user} ]]; then
			attn "gh must be authenticated to use this script."
            exists gh || ( gh --help; return 1; )
			gh auth login
            (( $? )) && ( attn "error running 'gh auth login' ... check https://cli.github.com/ "; return 1; )
            _gh_auth_username
            [[ -z ${_gh_user} ]] && ( attn "cannot log in to GitHub"; return 1; )
		fi

        dbecho "\$_gh_user (GitHub authenticated user): ${_gh_user}"
        dbecho "\${PWD} (current directory): ${PWD}"
    }
    _setup_local() {
        dbinfo "Setup local repo"

        #* check and setup local repo directory
        # create if needed and CD if possible
        if [[ -n "$1" ]]; then
            mkdir -p "$1" >/dev/null 2>&1
            cd "$1" || ( warn "error creating directory $1"; return 1 )
        fi

        # directory must be empty (certain parts of this setup can be run on existing repos)
        [ -n "$(ls -A ${PWD})" ] && ( warn "directory not empty"; return 1; )

    	#* Initial repo setup
			git init
            dbecho "\$?: $? - git init"
    }
    _setup_remote() {
        #* create remote repo from template (I use GitHub ... change it if you want)
            dbecho "gh repo create ${REPO_NAME} -y --public --template $GITHUB_TEMPLATE_PATH"
            gh repo create ${REPO_NAME} -y --public --template $GITHUB_TEMPLATE_PATH

            if (( $? )); then
                warn "error creating GitHub remote repo ${REPO_NAME}"
                is_prod && return 1; # dev mode may continue with existing repo
            fi

            # this should be done by gh ... but it doesn't always work
            dbecho git remote add origin "${GITHUB_REPO_URL}"
            git remote add origin "${GITHUB_REPO_URL}" >/dev/null 2>&1

            if (( $? )); then
                warn "error adding remote repository";
                is_prod && return 1; # dev mode may continue with existing repo
            fi

            # in the case of using an existing remote ... dev only
            dbecho git pull origin main --rebase
            git pull origin main --rebase >/dev/null 2>&1

            if (( $? )); then
                warn "error syncing remote repository"
                is_prod && return 1; # dev mode may continue with existing repo
            fi

        #* .gitignore and initial commit
            makeGI
            git add .gitignore -f
            git commit -m "initial commit"

            if (( $? )); then
                warn "error with initial commit";
                is_prod && return 1; # dev mode may continue with existing repo
            fi

            # push initial repository changes
            dbecho git push --set-upstream origin main
            git push --set-upstream origin main


            if (( $? )); then
                warn "error with initial remote repo push";
                is_prod && return 1; # dev mode may continue with existing repo
            fi
    }
    _setup_dirs() {
        # based on the unofficial and evolving https://github.com/golang-standards/project-layout

        # remove template placeholder example files
        rm -rf go.sum
        rm -rf go.mod
        rm -rf go.doc
        rm -rf gorepotemplate.go

        rm -rf cmd  # could use 'mv' here, but ... better to start fresh

        mkdir -p "$EXAMPLE_PATH"
        touch "$EXAMPLE_FILE"

        #* .gitignore and initial commit
        git add --all && git commit -m 'GoBot: setup directory tree and remove template examples'
        git push origin main
    }
    _make_files() {
		#* GitHub repo files
        GO_VERSION=$(_go_version)

        template gorepotemplate.go
        mv gorepotemplate.go ${REPO_NAME}.go

        # the default list of files is for template is:
            # CODE_OF_CONDUCT.md LICENSE README.md SECURITY.md contributing.md
            # go.doc go.test.sh idea.md .gitignore .editorconfig
        template

        git add --all && git commit -m "GoBot: setup repo files from template"

		#* GitHub Pages site setup
			git checkout -b gh-pages
            git fetch
			git push origin gh-pages

            git checkout main
			mkdir -p docs
            cd docs
            template docs/*
            cd ..

            git add --all && git commit -m "GoBot: create GitHub Pages folder branch"

        #* dev branch and initial dev version
            git checkout -b dev
            git fetch
            git push origin dev

            git checkout main
            mkdir -p .github
            cd .github
            template .github/*
            cd ISSUE_TEMPLATE
            template .github/ISSUE_TEMPLATE/*
            cd ..
            cd workflows
            template .github/workflows/*
            cd ${REPO_PATH}/${REPO_NAME}

            git add --all && devtag "GoBot: dev branch and directory tree setup"

		    #* Go module setup
            go mod init
            go get -u "${GO_GET_URL}"
            go mod tidy

            git add --all && git commit -m "GoBot: Go module setup"

            go doc >|go.doc
            chmod +x go.test.sh
            ./go.test.sh

            git add --all && git commit -m "GoBot: initial docs and test run"

    }

#? -----------------------------> main
    _gomake() {
        _setup_variables "$@"
        _setup_environment "$@"
        _setup_local "$@"
        _setup_remote "$@"
        _setup_dirs "$@"
        _make_files "$@"
        devtag "$@"
    }
    gomake() {
        case "$1" in

            -v|--version|version)
                echo "gomake version "
                return 0
                ;;

            -h|--help|help)
                echo "Usage: gomake [reponame] [--files] [--up]"
                return 0
                ;;

            --files|files)
                shift
                _make_files "$@"
                return
                ;;

            --devtag|devtag)
                shift
                devtag "$@"
                return
                ;;

            --bump|bump)
                shift
                version "$@"
                return
                ;;

            init)
                shift
                _gomake "$@"
                ;;

            *)

                usage gomake '[init|bump|help|files|devtag]'
                return 0
                ;;
        esac
    }

    #! repo testing ...
    # alias streamtest='cdgo; del stream; mkd stream; gomake'
    version=$(version)
    _setup_variables
    gomake "$@"
