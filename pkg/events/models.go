package events

import "time"

type GitlabEventTypeDetector struct {
	ObjectKind string `json:"object_kind"`
}

type GitlabPushEvent struct {
	ObjectKind   string      `json:"object_kind"`
	EventName    string      `json:"event_name"`
	Before       string      `json:"before"`
	After        string      `json:"after"`
	Ref          string      `json:"ref"`
	CheckoutSha  string      `json:"checkout_sha"`
	Message      interface{} `json:"message"`
	UserID       int         `json:"user_id"`
	UserName     string      `json:"user_name"`
	UserUsername string      `json:"user_username"`
	UserEmail    string      `json:"user_email"`
	UserAvatar   string      `json:"user_avatar"`
	ProjectID    int         `json:"project_id"`
	Project      struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	Commits []struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Added    []interface{} `json:"added"`
		Modified []string      `json:"modified"`
		Removed  []interface{} `json:"removed"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
	PushOptions       struct {
	} `json:"push_options"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
}

type GitlabTagPushEvent struct {
	ObjectKind   string      `json:"object_kind"`
	EventName    string      `json:"event_name"`
	Before       string      `json:"before"`
	After        string      `json:"after"`
	Ref          string      `json:"ref"`
	CheckoutSha  string      `json:"checkout_sha"`
	Message      interface{} `json:"message"`
	UserID       int         `json:"user_id"`
	UserName     string      `json:"user_name"`
	UserUsername string      `json:"user_username"`
	UserEmail    string      `json:"user_email"`
	UserAvatar   string      `json:"user_avatar"`
	ProjectID    int         `json:"project_id"`
	Project      struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	Commits []struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Added    []interface{} `json:"added"`
		Modified []string      `json:"modified"`
		Removed  []interface{} `json:"removed"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
	PushOptions       struct {
	} `json:"push_options"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
}

type GitlabNoteEvent struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	ProjectID int `json:"project_id"`
	Project   struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		Attachment       interface{} `json:"attachment"`
		AuthorID         int         `json:"author_id"`
		ChangePosition   interface{} `json:"change_position"`
		CommitID         string      `json:"commit_id"`
		CreatedAt        time.Time   `json:"created_at"`
		DiscussionID     string      `json:"discussion_id"`
		ID               int         `json:"id"`
		LineCode         interface{} `json:"line_code"`
		Note             string      `json:"note"`
		NoteableID       interface{} `json:"noteable_id"`
		NoteableType     string      `json:"noteable_type"`
		OriginalPosition interface{} `json:"original_position"`
		Position         interface{} `json:"position"`
		ProjectID        int         `json:"project_id"`
		ResolvedAt       interface{} `json:"resolved_at"`
		ResolvedByID     interface{} `json:"resolved_by_id"`
		ResolvedByPush   interface{} `json:"resolved_by_push"`
		StDiff           interface{} `json:"st_diff"`
		System           bool        `json:"system"`
		Type             interface{} `json:"type"`
		UpdatedAt        time.Time   `json:"updated_at"`
		UpdatedByID      interface{} `json:"updated_by_id"`
		Description      string      `json:"description"`
		URL              string      `json:"url"`
	} `json:"object_attributes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	Commit struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
}

type AutoGenerated struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		AssigneeID     interface{} `json:"assignee_id"`
		AuthorID       int         `json:"author_id"`
		CreatedAt      time.Time   `json:"created_at"`
		Description    string      `json:"description"`
		HeadPipelineID interface{} `json:"head_pipeline_id"`
		ID             int         `json:"id"`
		Iid            int         `json:"iid"`
		LastEditedAt   interface{} `json:"last_edited_at"`
		LastEditedByID interface{} `json:"last_edited_by_id"`
		MergeCommitSha interface{} `json:"merge_commit_sha"`
		MergeError     interface{} `json:"merge_error"`
		MergeParams    struct {
			ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
		} `json:"merge_params"`
		MergeStatus               string      `json:"merge_status"`
		MergeUserID               interface{} `json:"merge_user_id"`
		MergeWhenPipelineSucceeds bool        `json:"merge_when_pipeline_succeeds"`
		MilestoneID               interface{} `json:"milestone_id"`
		SourceBranch              string      `json:"source_branch"`
		SourceProjectID           int         `json:"source_project_id"`
		StateID                   int         `json:"state_id"`
		TargetBranch              string      `json:"target_branch"`
		TargetProjectID           int         `json:"target_project_id"`
		TimeEstimate              int         `json:"time_estimate"`
		Title                     string      `json:"title"`
		UpdatedAt                 time.Time   `json:"updated_at"`
		UpdatedByID               interface{} `json:"updated_by_id"`
		URL                       string      `json:"url"`
		Source                    struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      string      `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"source"`
		Target struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      string      `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID        string    `json:"id"`
			Message   string    `json:"message"`
			Title     string    `json:"title"`
			Timestamp time.Time `json:"timestamp"`
			URL       string    `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress      bool          `json:"work_in_progress"`
		TotalTimeSpent      int           `json:"total_time_spent"`
		HumanTotalTimeSpent interface{}   `json:"human_total_time_spent"`
		HumanTimeEstimate   interface{}   `json:"human_time_estimate"`
		AssigneeIds         []interface{} `json:"assignee_ids"`
		State               string        `json:"state"`
	} `json:"object_attributes"`
	Labels  []interface{} `json:"labels"`
	Changes struct {
	} `json:"changes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
}

type GitlabMergeRequestEvent struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		AssigneeID     interface{} `json:"assignee_id"`
		AuthorID       int         `json:"author_id"`
		CreatedAt      time.Time   `json:"created_at"`
		Description    string      `json:"description"`
		HeadPipelineID interface{} `json:"head_pipeline_id"`
		ID             int         `json:"id"`
		Iid            int         `json:"iid"`
		LastEditedAt   interface{} `json:"last_edited_at"`
		LastEditedByID interface{} `json:"last_edited_by_id"`
		MergeCommitSha interface{} `json:"merge_commit_sha"`
		MergeError     interface{} `json:"merge_error"`
		MergeParams    struct {
			ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
		} `json:"merge_params"`
		MergeStatus               string      `json:"merge_status"`
		MergeUserID               interface{} `json:"merge_user_id"`
		MergeWhenPipelineSucceeds bool        `json:"merge_when_pipeline_succeeds"`
		MilestoneID               interface{} `json:"milestone_id"`
		SourceBranch              string      `json:"source_branch"`
		SourceProjectID           int         `json:"source_project_id"`
		StateID                   int         `json:"state_id"`
		TargetBranch              string      `json:"target_branch"`
		TargetProjectID           int         `json:"target_project_id"`
		TimeEstimate              int         `json:"time_estimate"`
		Title                     string      `json:"title"`
		UpdatedAt                 time.Time   `json:"updated_at"`
		UpdatedByID               interface{} `json:"updated_by_id"`
		URL                       string      `json:"url"`
		Source                    struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      string      `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"source"`
		Target struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      string      `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID        string    `json:"id"`
			Message   string    `json:"message"`
			Title     string    `json:"title"`
			Timestamp time.Time `json:"timestamp"`
			URL       string    `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress      bool          `json:"work_in_progress"`
		TotalTimeSpent      int           `json:"total_time_spent"`
		HumanTotalTimeSpent interface{}   `json:"human_total_time_spent"`
		HumanTimeEstimate   interface{}   `json:"human_time_estimate"`
		AssigneeIds         []interface{} `json:"assignee_ids"`
		State               string        `json:"state"`
	} `json:"object_attributes"`
	Labels  []interface{} `json:"labels"`
	Changes struct {
	} `json:"changes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
}

type GitlabJobEvent struct {
	ObjectKind         string      `json:"object_kind"`
	Ref                string      `json:"ref"`
	Tag                bool        `json:"tag"`
	BeforeSha          string      `json:"before_sha"`
	Sha                string      `json:"sha"`
	BuildID            int         `json:"build_id"`
	BuildName          string      `json:"build_name"`
	BuildStage         string      `json:"build_stage"`
	BuildStatus        string      `json:"build_status"`
	BuildStartedAt     time.Time   `json:"build_started_at"`
	BuildFinishedAt    time.Time   `json:"build_finished_at"`
	BuildDuration      float64     `json:"build_duration"`
	BuildAllowFailure  bool        `json:"build_allow_failure"`
	BuildFailureReason string      `json:"build_failure_reason"`
	PipelineID         int         `json:"pipeline_id"`
	Runner             interface{} `json:"runner"`
	ProjectID          int         `json:"project_id"`
	ProjectName        string      `json:"project_name"`
	User               struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Commit struct {
		ID          int       `json:"id"`
		Sha         string    `json:"sha"`
		Message     string    `json:"message"`
		AuthorName  string    `json:"author_name"`
		AuthorEmail string    `json:"author_email"`
		AuthorURL   string    `json:"author_url"`
		Status      string    `json:"status"`
		Duration    int       `json:"duration"`
		StartedAt   time.Time `json:"started_at"`
		FinishedAt  time.Time `json:"finished_at"`
	} `json:"commit"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
}

type GitlabPipelineEvent struct {
	ObjectKind       string `json:"object_kind"`
	ObjectAttributes struct {
		ID             int           `json:"id"`
		Ref            string        `json:"ref"`
		Tag            bool          `json:"tag"`
		Sha            string        `json:"sha"`
		BeforeSha      string        `json:"before_sha"`
		Source         string        `json:"source"`
		Status         string        `json:"status"`
		DetailedStatus string        `json:"detailed_status"`
		Stages         []string      `json:"stages"`
		CreatedAt      time.Time     `json:"created_at"`
		FinishedAt     time.Time     `json:"finished_at"`
		Duration       int           `json:"duration"`
		Variables      []interface{} `json:"variables"`
	} `json:"object_attributes"`
	MergeRequest interface{} `json:"merge_request"`
	User         struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
	} `json:"project"`
	Commit struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
	Builds []struct {
		ID           int       `json:"id"`
		Stage        string    `json:"stage"`
		Name         string    `json:"name"`
		Status       string    `json:"status"`
		CreatedAt    time.Time `json:"created_at"`
		StartedAt    time.Time `json:"started_at"`
		FinishedAt   time.Time `json:"finished_at"`
		When         string    `json:"when"`
		Manual       bool      `json:"manual"`
		AllowFailure bool      `json:"allow_failure"`
		User         struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			AvatarURL string `json:"avatar_url"`
			Email     string `json:"email"`
		} `json:"user"`
		Runner        interface{} `json:"runner"`
		ArtifactsFile struct {
			Filename interface{} `json:"filename"`
			Size     interface{} `json:"size"`
		} `json:"artifacts_file"`
	} `json:"builds"`
}
