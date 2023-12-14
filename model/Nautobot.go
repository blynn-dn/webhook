package model

import "time"

// WebhookEvent The following was produced with an online helper tool (https://transform.tools/json-to-go) was used to
// generate the struct from JSON.  The generated struct is used as-is.
//
// Note that there are many redundant/repetitive entries.
// For example; sub-structs Snapshots.Prechange, .PostChange and .Differences, have the exact or
// very near the same structures.  Data contains much of the same structure as the aforementioned sub-struct.
// A bit of refactoring to normalized the entries structure is in order.
//
// todo should normalize WebhookEvent
type WebhookEvent struct {
    Event     string `json:"event"`
    Timestamp string `json:"timestamp"`
    Model     string `json:"model"`
    Username  string `json:"username"`
    RequestID string `json:"request_id"`
    Data      struct {
        ID         string `json:"id"`
        Display    string `json:"display"`
        URL        string `json:"url"`
        Name       string `json:"name"`
        DeviceType struct {
            Display      string `json:"display"`
            ID           string `json:"id"`
            URL          string `json:"url"`
            Manufacturer struct {
                Display string `json:"display"`
                ID      string `json:"id"`
                URL     string `json:"url"`
                Name    string `json:"name"`
                Slug    string `json:"slug"`
            } `json:"manufacturer"`
            Model string `json:"model"`
            Slug  string `json:"slug"`
        } `json:"device_type"`
        DeviceRole struct {
            Display string `json:"display"`
            ID      string `json:"id"`
            URL     string `json:"url"`
            Name    string `json:"name"`
            Slug    string `json:"slug"`
        } `json:"device_role"`
        Tenant   interface{} `json:"tenant"`
        Platform interface{} `json:"platform"`
        Serial   string      `json:"serial"`
        AssetTag interface{} `json:"asset_tag"`
        Site     struct {
            Display string `json:"display"`
            ID      string `json:"id"`
            URL     string `json:"url"`
            Name    string `json:"name"`
            Slug    string `json:"slug"`
        } `json:"site"`
        Location struct {
            Display   string `json:"display"`
            ID        string `json:"id"`
            URL       string `json:"url"`
            Name      string `json:"name"`
            Slug      string `json:"slug"`
            TreeDepth int    `json:"tree_depth"`
        } `json:"location"`
        Rack struct {
            Display string `json:"display"`
            ID      string `json:"id"`
            URL     string `json:"url"`
            Name    string `json:"name"`
        } `json:"rack"`
        Position     interface{} `json:"position"`
        Face         interface{} `json:"face"`
        ParentDevice interface{} `json:"parent_device"`
        Status       struct {
            Value string `json:"value"`
            Label string `json:"label"`
        } `json:"status"`
        PrimaryIP struct {
            Display string `json:"display"`
            ID      string `json:"id"`
            URL     string `json:"url"`
            Family  int    `json:"family"`
            Address string `json:"address"`
        } `json:"primary_ip"`
        PrimaryIP4 struct {
            Display string `json:"display"`
            ID      string `json:"id"`
            URL     string `json:"url"`
            Family  int    `json:"family"`
            Address string `json:"address"`
        } `json:"primary_ip4"`
        PrimaryIP6                    interface{} `json:"primary_ip6"`
        SecretsGroup                  interface{} `json:"secrets_group"`
        Cluster                       interface{} `json:"cluster"`
        VirtualChassis                interface{} `json:"virtual_chassis"`
        VcPosition                    interface{} `json:"vc_position"`
        VcPriority                    interface{} `json:"vc_priority"`
        DeviceRedundancyGroup         interface{} `json:"device_redundancy_group"`
        DeviceRedundancyGroupPriority interface{} `json:"device_redundancy_group_priority"`
        Comments                      string      `json:"comments"`
        LocalContextSchema            interface{} `json:"local_context_schema"`
        LocalContextData              interface{} `json:"local_context_data"`
        Created                       string      `json:"created"`
        LastUpdated                   time.Time   `json:"last_updated"`
        Tags                          []struct {
            Display     string    `json:"display"`
            ID          string    `json:"id"`
            URL         string    `json:"url"`
            Name        string    `json:"name"`
            Slug        string    `json:"slug"`
            Color       string    `json:"color"`
            Created     string    `json:"created"`
            LastUpdated time.Time `json:"last_updated"`
        } `json:"tags"`
        NotesURL     string `json:"notes_url"`
        CustomFields struct {
        } `json:"custom_fields"`
    } `json:"data"`
    Snapshots struct {
        Prechange struct {
            ID   string      `json:"id"`
            URL  string      `json:"url"`
            Face interface{} `json:"face"`
            Name string      `json:"name"`
            Rack struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Name    string `json:"name"`
                Display string `json:"display"`
            } `json:"rack"`
            Site struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Name    string `json:"name"`
                Slug    string `json:"slug"`
                Display string `json:"display"`
            } `json:"site"`
            Tags []struct {
                ID          string    `json:"id"`
                URL         string    `json:"url"`
                Name        string    `json:"name"`
                Slug        string    `json:"slug"`
                Color       string    `json:"color"`
                Created     string    `json:"created"`
                Display     string    `json:"display"`
                LastUpdated time.Time `json:"last_updated"`
            } `json:"tags"`
            Serial string `json:"serial"`
            Status struct {
                Label string `json:"label"`
                Value string `json:"value"`
            } `json:"status"`
            Tenant   interface{} `json:"tenant"`
            Cluster  interface{} `json:"cluster"`
            Created  string      `json:"created"`
            Display  string      `json:"display"`
            Comments string      `json:"comments"`
            Location struct {
                ID        string `json:"id"`
                URL       string `json:"url"`
                Name      string `json:"name"`
                Slug      string `json:"slug"`
                Display   string `json:"display"`
                TreeDepth int    `json:"tree_depth"`
            } `json:"location"`
            Platform  interface{} `json:"platform"`
            Position  interface{} `json:"position"`
            AssetTag  interface{} `json:"asset_tag"`
            NotesURL  string      `json:"notes_url"`
            PrimaryIP struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Family  int    `json:"family"`
                Address string `json:"address"`
                Display string `json:"display"`
            } `json:"primary_ip"`
            DeviceRole struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Name    string `json:"name"`
                Slug    string `json:"slug"`
                Display string `json:"display"`
            } `json:"device_role"`
            DeviceType struct {
                ID           string `json:"id"`
                URL          string `json:"url"`
                Slug         string `json:"slug"`
                Model        string `json:"model"`
                Display      string `json:"display"`
                Manufacturer struct {
                    ID      string `json:"id"`
                    URL     string `json:"url"`
                    Name    string `json:"name"`
                    Slug    string `json:"slug"`
                    Display string `json:"display"`
                } `json:"manufacturer"`
            } `json:"device_type"`
            PrimaryIP4 struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Family  int    `json:"family"`
                Address string `json:"address"`
                Display string `json:"display"`
            } `json:"primary_ip4"`
            PrimaryIP6   interface{} `json:"primary_ip6"`
            VcPosition   interface{} `json:"vc_position"`
            VcPriority   interface{} `json:"vc_priority"`
            LastUpdated  time.Time   `json:"last_updated"`
            CustomFields struct {
            } `json:"custom_fields"`
            ParentDevice                  interface{} `json:"parent_device"`
            SecretsGroup                  interface{} `json:"secrets_group"`
            VirtualChassis                interface{} `json:"virtual_chassis"`
            LocalContextData              interface{} `json:"local_context_data"`
            LocalContextSchema            interface{} `json:"local_context_schema"`
            DeviceRedundancyGroup         interface{} `json:"device_redundancy_group"`
            DeviceRedundancyGroupPriority interface{} `json:"device_redundancy_group_priority"`
        } `json:"prechange"`
        Postchange struct {
            ID   string      `json:"id"`
            URL  string      `json:"url"`
            Face interface{} `json:"face"`
            Name string      `json:"name"`
            Rack struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Name    string `json:"name"`
                Display string `json:"display"`
            } `json:"rack"`
            Site struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Name    string `json:"name"`
                Slug    string `json:"slug"`
                Display string `json:"display"`
            } `json:"site"`
            Tags []struct {
                ID          string    `json:"id"`
                URL         string    `json:"url"`
                Name        string    `json:"name"`
                Slug        string    `json:"slug"`
                Color       string    `json:"color"`
                Created     string    `json:"created"`
                Display     string    `json:"display"`
                LastUpdated time.Time `json:"last_updated"`
            } `json:"tags"`
            Serial string `json:"serial"`
            Status struct {
                Label string `json:"label"`
                Value string `json:"value"`
            } `json:"status"`
            Tenant   interface{} `json:"tenant"`
            Cluster  interface{} `json:"cluster"`
            Created  string      `json:"created"`
            Display  string      `json:"display"`
            Comments string      `json:"comments"`
            Location struct {
                ID        string `json:"id"`
                URL       string `json:"url"`
                Name      string `json:"name"`
                Slug      string `json:"slug"`
                Display   string `json:"display"`
                TreeDepth int    `json:"tree_depth"`
            } `json:"location"`
            Platform  interface{} `json:"platform"`
            Position  interface{} `json:"position"`
            AssetTag  interface{} `json:"asset_tag"`
            NotesURL  string      `json:"notes_url"`
            PrimaryIP struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Family  int    `json:"family"`
                Address string `json:"address"`
                Display string `json:"display"`
            } `json:"primary_ip"`
            DeviceRole struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Name    string `json:"name"`
                Slug    string `json:"slug"`
                Display string `json:"display"`
            } `json:"device_role"`
            DeviceType struct {
                ID           string `json:"id"`
                URL          string `json:"url"`
                Slug         string `json:"slug"`
                Model        string `json:"model"`
                Display      string `json:"display"`
                Manufacturer struct {
                    ID      string `json:"id"`
                    URL     string `json:"url"`
                    Name    string `json:"name"`
                    Slug    string `json:"slug"`
                    Display string `json:"display"`
                } `json:"manufacturer"`
            } `json:"device_type"`
            PrimaryIP4 struct {
                ID      string `json:"id"`
                URL     string `json:"url"`
                Family  int    `json:"family"`
                Address string `json:"address"`
                Display string `json:"display"`
            } `json:"primary_ip4"`
            PrimaryIP6   interface{} `json:"primary_ip6"`
            VcPosition   interface{} `json:"vc_position"`
            VcPriority   interface{} `json:"vc_priority"`
            LastUpdated  time.Time   `json:"last_updated"`
            CustomFields struct {
            } `json:"custom_fields"`
            ParentDevice                  interface{} `json:"parent_device"`
            SecretsGroup                  interface{} `json:"secrets_group"`
            VirtualChassis                interface{} `json:"virtual_chassis"`
            LocalContextData              interface{} `json:"local_context_data"`
            LocalContextSchema            interface{} `json:"local_context_schema"`
            DeviceRedundancyGroup         interface{} `json:"device_redundancy_group"`
            DeviceRedundancyGroupPriority interface{} `json:"device_redundancy_group_priority"`
        } `json:"postchange"`
        Differences struct {
            Removed struct {
                PrimaryIP struct {
                    ID      string `json:"id"`
                    URL     string `json:"url"`
                    Family  int    `json:"family"`
                    Address string `json:"address"`
                    Display string `json:"display"`
                } `json:"primary_ip"`
                PrimaryIP4 struct {
                    ID      string `json:"id"`
                    URL     string `json:"url"`
                    Family  int    `json:"family"`
                    Address string `json:"address"`
                    Display string `json:"display"`
                } `json:"primary_ip4"`
            } `json:"removed"`
            Added struct {
                PrimaryIP struct {
                    ID      string `json:"id"`
                    URL     string `json:"url"`
                    Family  int    `json:"family"`
                    Address string `json:"address"`
                    Display string `json:"display"`
                } `json:"primary_ip"`
                PrimaryIP4 struct {
                    ID      string `json:"id"`
                    URL     string `json:"url"`
                    Family  int    `json:"family"`
                    Address string `json:"address"`
                    Display string `json:"display"`
                } `json:"primary_ip4"`
            } `json:"added"`
        } `json:"differences"`
    } `json:"snapshots"`
}
